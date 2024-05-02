package controller

import (
	"bookstore/cmd/dao"
	"bookstore/cmd/model"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func CheckoutHandler(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLoggedIn(r)
	userID := session.UserID

	cart, err := dao.GetCartByUserID(userID)
	if err != nil {
		log.Printf("GetCartByUserID failed: %v", err)
		http.Error(w, "Failed to get cart by user ID", http.StatusInternalServerError)
		return
	}

	orderID := uuid.NewString()
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  cart.TotalCount,
		TotalAmount: cart.TotalAmount,
		Status:      0,
		UserID:      userID,
	}

	if err := dao.AddOrder(order); err != nil {
		log.Printf("AddOrder failed: %v", err)
		http.Error(w, "Failed to add order", http.StatusInternalServerError)
		return
	}

	cartItems := cart.CartItems
	for _, v := range cartItems {
		orderItem := &model.OrderItem{
			Count:   v.Count,
			Amount:  v.Amount,
			Title:   v.Book.Title,
			Author:  v.Book.Author,
			Price:   v.Book.Price,
			ImgPath: v.Book.ImgPath,
			OrderID: orderID,
		}

		if err := dao.AddOrderItem(orderItem); err != nil {
			log.Printf("AddOrderItem failed: %v", err)
			http.Error(w, "Failed to add order item", http.StatusInternalServerError)
			return
		}

		book := v.Book
		book.Sales += v.Count
		book.Stock -= v.Count
		if err := dao.UpdateBook(book); err != nil {
			log.Printf("UpdateBook failed: %v", err)
			http.Error(w, "Failed to update book", http.StatusInternalServerError)
			return
		}
	}

	if err := dao.DeleteCartByCartID(cart.CartID); err != nil {
		log.Printf("DeleteCartByCartID failed: %v", err)
		http.Error(w, "Failed to delete cart by cart ID", http.StatusInternalServerError)
		return
	}

	session.Order = order

	tmpl := template.Must(template.ParseFiles("views/pages/cart/checkout.html"))
	if err := tmpl.Execute(w, session); err != nil {
		log.Printf("Failed to execute CheckoutHandler template: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func GetOrdersHandler(w http.ResponseWriter, r *http.Request) {
	orders, _ := dao.GetOrders()

	tmpl := template.Must(template.ParseFiles("views/pages/order/order_manager.html"))
	if err := tmpl.Execute(w, orders); err != nil {
		log.Printf("Failed to execute GetOrdersHandler template: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func GetOrderInfoHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")
	orderItems, err := dao.GetOrderItemsByOrderID(orderID)
	if err != nil {
		log.Printf("GetOrderItemsByOrderID failed: %v", err)
		http.Error(w, "Failed to get order items by order ID", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("views/pages/order/order_info.html"))
	if err := tmpl.Execute(w, orderItems); err != nil {
		log.Printf("Failed to execute GetOrderInfoHandler template: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func GetMyOrdersHandler(w http.ResponseWriter, r *http.Request) {
	_, session := dao.IsLoggedIn(r)
	userID := session.UserID

	orders, _ := dao.GetMyOrders(userID)
	session.Orders = orders

	tmpl := template.Must(template.ParseFiles("views/pages/order/order.html"))
	if err := tmpl.Execute(w, session); err != nil {
		log.Printf("Failed to execute GetMyOrdersHandler template: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func ShipOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")

	if err := dao.UpdateOrderStatus(orderID, 1); err != nil {
		log.Printf("UpdateOrderStatus failed: %v", err)
		http.Error(w, "Failed to update order status", http.StatusInternalServerError)
		return
	}

	GetOrdersHandler(w, r)
}

func ReceiveOrderHandler(w http.ResponseWriter, r *http.Request) {
	orderID := r.FormValue("orderID")

	if err := dao.UpdateOrderStatus(orderID, 2); err != nil {
		log.Printf("ReceiveOrder failed: %v", err)
		http.Error(w, "Failed to receive order", http.StatusInternalServerError)
		return
	}

	GetMyOrdersHandler(w, r)
}
