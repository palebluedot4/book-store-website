package controller

import (
	"bookstore/cmd/dao"
	"bookstore/cmd/model"
	"encoding/json"
	"html/template"
	"log"
	"net/http"
	"strconv"

	"github.com/google/uuid"
)

func AddToCartHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the user is logged
	loggedIn, session := dao.IsLoggedIn(r)
	if loggedIn { // Check if the user is logged in to allow adding items to the cart
		bookID := r.FormValue("bookID")

		// Retrieve information about the book user wants to purchase
		book, err := dao.GetBookByID(bookID)
		if err != nil {
			log.Printf("GetBookByID failed: %v", err)
			http.Error(w, "Failed to retrieve book", http.StatusInternalServerError)
			return
		}

		userID := session.UserID

		// Get the user's cart, automatically creating a new cart if it doesn't exist
		// Logging errors due to the absence of a cart is unnecessary since a new cart will be created if none exists
		cart, _ := dao.GetCartByUserID(userID)

		switch {
		case cart != nil: // If the user has a cart
			// Check if the book already exists in the cart
			// Logging errors due to the absence of a cart is unnecessary since a new cart will be created if none exists
			cartItem, _ := dao.GetCartItemByBookIDAndCartID(bookID, cart.CartID)

			if cartItem != nil { // If the book is already in the cart
				// Update the quantity of the book in the cart
				cartItems := cart.CartItems
				for _, v := range cartItems {
					if v.Book.ID == cartItem.Book.ID {
						v.Count += 1
						if err := dao.UpdateCartItemQuantity(v); err != nil {
							log.Printf("UpdateCartItemQuantity failed: %v", err)
							http.Error(w, "Failed update cart item quantity", http.StatusInternalServerError)
							return
						}
					}
				}
			} else { // If the book is not in the cart
				// Create a new cart item and add it to the cart
				cartItem := &model.CartItem{
					Book:   book,
					Count:  1,
					CartID: cart.CartID,
				}

				cart.CartItems = append(cart.CartItems, cartItem)
				if err := dao.AddCartItem(cartItem); err != nil {
					log.Printf("AddCartItem failed: %v", err)
					http.Error(w, "Failed to add cart item", http.StatusInternalServerError)
					return
				}
			}

			// Ensure the database's cart stays updated with any handling process changes
			if err := dao.UpdateCart(cart); err != nil {
				log.Printf("UpdateCart failed: %v", err)
				http.Error(w, "Failed to update cart", http.StatusInternalServerError)
				return
			}

		default: // If the user does not have a cart, create a new one
			cartID := uuid.NewString()
			cart := &model.Cart{
				CartID: cartID,
				UserID: userID,
			}

			var cartItems []*model.CartItem
			cartItem := &model.CartItem{
				Book:   book,
				Count:  1,
				CartID: cartID,
			}
			cartItems = append(cartItems, cartItem)
			cart.CartItems = cartItems

			if err := dao.AddCart(cart); err != nil {
				log.Printf("AddCart failed: %v", err)
				http.Error(w, "Failed to add cart", http.StatusInternalServerError)
				return
			}
		}

		// Display the title of the book that has been added to the user's cart on the frontend
		if _, err := w.Write([]byte("商品" + book.Title + "已加入購物車中")); err != nil {
			log.Printf("Failed to write response: %v", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	} else { // If not logged in, redirect the user to the login page
		if _, err := w.Write([]byte("To login page")); err != nil {
			log.Printf("Failed to write response: %v", err)
			http.Error(w, "Failed to write response", http.StatusInternalServerError)
			return
		}
	}
}

func GetCartInfoHandler(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLoggedIn(r)
	userID := session.UserID
	cart, _ := dao.GetCartByUserID(userID)

	switch {
	case cart != nil:
		session.Cart = cart
		tmpl := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		if err := tmpl.Execute(w, session); err != nil {
			log.Printf("Failed to execute GetCartInfoHandler template: %s", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	default:
		tmpl := template.Must(template.ParseFiles("views/pages/cart/cart.html"))
		if err := tmpl.Execute(w, session); err != nil {
			log.Printf("Failed to execute GetCartInfoHandler template: %s", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	}
}

func DeleteCartHandler(w http.ResponseWriter, r *http.Request) {
	cartID := r.FormValue("cartID")

	if err := dao.DeleteCartByCartID(cartID); err != nil {
		log.Printf("DeleteCartByCartID failed: %v", err)
		http.Error(w, "Failed to delete cart", http.StatusInternalServerError)
		return
	}

	GetCartInfoHandler(w, r)
}

func DeleteCartItemHandler(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemID")
	intCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)

	_, session := dao.IsLoggedIn(r)
	userID := session.UserID

	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Printf("GetCartByUserID failed: %v", err)
		http.Error(w, "Failed to get cart by user ID", http.StatusInternalServerError)
		return
	}

	cartItems := cart.CartItems
	for i, v := range cartItems {
		if v.ID == intCartItemID {
			cartItems = append(cartItems[:i], cartItems[i+1:]...)
			cart.CartItems = cartItems

			if err := dao.DeleteCartItemByID(cartItemID); err != nil {
				log.Printf("DeleteCartItemByID failed: %v", err)
				http.Error(w, "Failed to delete cart item", http.StatusInternalServerError)
				return
			}
		}
	}

	if err := dao.UpdateCart(cart); err != nil {
		log.Printf("UpdateCart failed: %v", err)
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	GetCartInfoHandler(w, r)
}

func UpdateCartItemHandler(w http.ResponseWriter, r *http.Request) {
	cartItemID := r.FormValue("cartItemID")
	bookCount := r.FormValue("bookCount")
	intCartItemID, _ := strconv.ParseInt(cartItemID, 10, 64)
	intBookCount, _ := strconv.ParseInt(bookCount, 10, 64)

	_, session := dao.IsLoggedIn(r)
	userID := session.UserID

	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Printf("GetCartByUserID failed: %v", err)
		http.Error(w, "Failed to get cart by user ID", http.StatusInternalServerError)
		return
	}

	for _, v := range cart.CartItems {
		if v.ID == intCartItemID {
			v.Count = intBookCount
			if err := dao.UpdateCartItemQuantity(v); err != nil {
				log.Printf("UpdateCartItemQuantity failed: %v", err)
				http.Error(w, "Failed to update cart item quantity", http.StatusInternalServerError)
			}
		}
	}

	if err := dao.UpdateCart(cart); err != nil {
		log.Printf("UpdateCart failed: %v", err)
		http.Error(w, "Failed to update cart", http.StatusInternalServerError)
		return
	}

	// GetCartInfoHandler(w, r)
	updatedCart, _ := dao.GetCartByUserID(userID)

	var cartItemAmount float64
	totalCount := updatedCart.TotalCount
	totalAmount := updatedCart.TotalAmount

	for _, v := range updatedCart.CartItems {
		if v.ID == intCartItemID {
			cartItemAmount = v.Amount
		}
	}

	cartItemResponseData := model.CartItemResponse{
		Amount:      cartItemAmount,
		TotalCount:  totalCount,
		TotalAmount: totalAmount,
	}

	responseJSON, err := json.Marshal(cartItemResponseData)
	if err != nil {
		log.Printf("Error encoding JSON data: %v", err)
		http.Error(w, "Failed to encode JSON data", http.StatusInternalServerError)
		return
	}

	if _, err := w.Write(responseJSON); err != nil {
		log.Printf("Failed to write response: %v", err)
		http.Error(w, "Failed to write response", http.StatusInternalServerError)
		return
	}
}
