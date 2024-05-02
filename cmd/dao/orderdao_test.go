package dao

import (
	"bookstore/cmd/model"
	"fmt"
	"testing"
	"time"
)

func TestOrder(t *testing.T) {
	// t.Run("Add order validation", testAddOrder)
	// t.Run("Get orders validation", testGetOrders)
	// t.Run("Get my orders validation", testGetMyOrders)
	// t.Run("Update order status validation", testUpdateOrderStatus)
}

func testAddOrder(t *testing.T) {
	orderID := "test"
	order := &model.Order{
		OrderID:     orderID,
		CreateTime:  time.Now().Format("2006-01-02 15:04:05"),
		TotalCount:  2,
		TotalAmount: 720,
		Status:      0,
		UserID:      1,
	}
	orderItem1 := &model.OrderItem{
		Count:   1,
		Amount:  340,
		Title:   "腹語術",
		Author:  "夏宇",
		Price:   340,
		ImgPath: "static/img/default.jpg",
		OrderID: orderID,
	}
	orderItem2 := &model.OrderItem{
		Count:   1,
		Amount:  380,
		Title:   "詩60首",
		Author:  "夏宇",
		Price:   380,
		ImgPath: "static/img/default.jpg",
		OrderID: orderID,
	}

	if err := AddOrder(order); err != nil {
		t.Fatalf("AddOrder failed: %v", err)
	}
	if err := AddOrderItem(orderItem1); err != nil {
		t.Fatalf("AddOrderItem failed: %v", err)
	}
	if err := AddOrderItem(orderItem2); err != nil {
		t.Fatalf("AddOrderItem failed: %v", err)
	}
}

func testGetOrders(t *testing.T) {
	orders, err := GetOrders()
	if err != nil {
		t.Fatalf("GetOrders failed: %v", err)
	}

	for i, v := range orders {
		fmt.Printf("Information for orders #%d: %v\n", i+1, v)
	}
}

func testGetMyOrders(t *testing.T) {
	var userID int64 = 2
	orders, err := GetMyOrders(userID)
	if err != nil {
		t.Fatalf("GetMyOrders failed: %v", err)
	}

	for i, v := range orders {
		fmt.Printf("Information for my orders #%d: %v\n", i+1, v)
	}
}

func testUpdateOrderStatus(t *testing.T) {
	orderID := "c5eca9c8-59f1-4ebf-ad5b-5f390ed385e5"
	var status model.OrderStatus = 1

	err := UpdateOrderStatus(orderID, status)
	if err != nil {
		t.Fatalf("testUpdateOrderStatus failed: %v", err)
	}
}
