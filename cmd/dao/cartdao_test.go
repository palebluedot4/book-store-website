package dao

import (
	"bookstore/cmd/model"
	"fmt"
	"testing"
)

func TestCart(t *testing.T) {
	// t.Run("Add cart validation", testAddCart)
	// t.Run("Get cart by user ID validation", testGetCartByUserID)
	// t.Run("Update cart validation", testUpdateCart)
	// t.Run("Delete cart by cart ID validation", testDeleteCartByCartID)
}

func testAddCart(t *testing.T) {
	book1 := &model.Book{
		ID:    1,
		Price: 340.00,
	}
	book2 := &model.Book{
		ID:    2,
		Price: 380.00,
	}
	cartItem1 := &model.CartItem{
		Book:   book1,
		Count:  2,
		CartID: "test",
	}
	cartItem2 := &model.CartItem{
		Book:   book2,
		Count:  4,
		CartID: "test",
	}

	var cartItems []*model.CartItem
	cartItems = append(cartItems, cartItem1)
	cartItems = append(cartItems, cartItem2)
	cart := &model.Cart{
		CartID:    "test",
		CartItems: cartItems,
		UserID:    1,
	}

	if err := AddCart(cart); err != nil {
		t.Fatalf("AddCart failed: %v", err)
	}
}

func testGetCartByUserID(t *testing.T) {
	userID := 1
	cart, err := GetCartByUserID(int64(userID))
	if err != nil {
		t.Fatalf("GetCartByUserID failed: %v", err)
	}
	fmt.Printf("Information for cart: CartID: %s, CartItems: %v, TotalCount: %d, TotalAmount: %.2f, UserID: %d\n", cart.CartID, cart.CartItems, cart.TotalCount, cart.TotalAmount, cart.UserID)
}

func testUpdateCart(t *testing.T) {
	updatedCart := &model.Cart{
		CartID: "test",
		CartItems: []*model.CartItem{
			{
				ID:    1,
				Count: 10,
				Book: &model.Book{
					ID:    1,
					Price: 340.00,
				},
			},
			{
				ID:    2,
				Count: 10,
				Book: &model.Book{
					ID:    2,
					Price: 380.00,
				},
			},
		},
		UserID: 1,
	}

	if err := UpdateCart(updatedCart); err != nil {
		t.Fatalf("UpdateCart failed: %v", err)

		updatedCart, err := GetCartByUserID(1)
		if err != nil {
			t.Fatalf("GetCartByUserID failed: %v", err)
		}
		fmt.Printf("Cart updated successfully: CartID: %s, TotalCount: %d, TotalAmount: %.2f\n", updatedCart.CartID, updatedCart.TotalCount, updatedCart.TotalAmount)
	}
}

func testDeleteCartByCartID(t *testing.T) {
	cartID := "0507c7fe-6be0-46e9-ac87-48198f22c41c"

	if err := DeleteCartByCartID(cartID); err != nil {
		t.Fatalf("DeleteCartByCartID failed: %v", err)
	}
}
