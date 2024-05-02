package main

import (
	"bookstore/cmd/controller"
	"log"
	"net/http"
)

func main() {
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	// http.HandleFunc("/main", controller.IndexHandler)
	http.HandleFunc("/main", controller.FetchBooksByPriceRangeHandler)

	http.HandleFunc("/login", controller.LoginHandler)
	http.HandleFunc("/logout", controller.LogoutHandler)
	http.HandleFunc("/register", controller.RegisterHandler)
	http.HandleFunc("/checkUserName", controller.CheckUserNameHandler)

	// http.HandleFunc("/getBooks", controller.GetBooksHandler)
	// http.HandleFunc("/addBook", controller.AddBookHandler)
	http.HandleFunc("/deleteBook", controller.DeleteBookHandler)
	http.HandleFunc("/editBook", controller.EditBookHandler)
	// http.HandleFunc("/updateBook", controller.UpdateBookHandler)
	http.HandleFunc("/updateOrAddBook", controller.UpdateOrAddBookHandler)
	http.HandleFunc("/getPaginatedBooks", controller.GetPaginatedBooksHandler)
	http.HandleFunc("/fetchBooksByPriceRange", controller.FetchBooksByPriceRangeHandler)

	http.HandleFunc("/addToCart", controller.AddToCartHandler)
	http.HandleFunc("/getCartInfo", controller.GetCartInfoHandler)
	http.HandleFunc("/deleteCart", controller.DeleteCartHandler)
	http.HandleFunc("/deleteCartItem", controller.DeleteCartItemHandler)
	http.HandleFunc("/updateCartItem", controller.UpdateCartItemHandler)

	http.HandleFunc("/checkout", controller.CheckoutHandler)
	http.HandleFunc("/getOrders", controller.GetOrdersHandler)
	http.HandleFunc("/getOrderInfo", controller.GetOrderInfoHandler)
	http.HandleFunc("/getMyOrders", controller.GetMyOrdersHandler)
	http.HandleFunc("/shipOrder", controller.ShipOrderHandler)
	http.HandleFunc("/receiveOrder", controller.ReceiveOrderHandler)

	log.Println("Listening and serving HTTP on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
