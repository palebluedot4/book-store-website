package dao

import (
	"fmt"
	"testing"

	"bookstore/cmd/model"
)

func TestOrderItem(t *testing.T) {
	t.Run("Add order item validation", testAddOrderItem)
	t.Run("Get order items by order ID validation", testGetOrderItemsByOrderID)
}

func testAddOrderItem(t *testing.T) {
	orderID := "test"
	orderItem := &model.OrderItem{
		Count:   1,
		Amount:  340,
		Title:   "腹語術",
		Author:  "夏宇",
		Price:   340,
		ImgPath: "static/img/default.jpg",
		OrderID: orderID,
	}

	if err := AddOrderItem(orderItem); err != nil {
		t.Fatalf("AddOrderItem failed: %v", err)
	}
}

func testGetOrderItemsByOrderID(t *testing.T) {
	orderID := "c5eca9c8-59f1-4ebf-ad5b-5f390ed385e5"
	orderItems, err := GetOrderItemsByOrderID(orderID)
	if err != nil {
		t.Fatalf("GetOrderItemsByOrderID failed: %v", err)
	}

	for i, v := range orderItems {
		fmt.Printf("Information for order items #%d: %v\n", i+1, v)
	}
}
