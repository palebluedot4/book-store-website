package dao

import (
	"fmt"

	"bookstore/cmd/model"
	"bookstore/cmd/utils"
)

func AddOrderItem(orderItem *model.OrderItem) error {
	const query = "INSERT INTO order_items (count, amount, title, author, price, img_path, order_id) VALUES (?, ?, ?, ?, ?, ?, ?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(orderItem.Count, orderItem.Amount, orderItem.Title, orderItem.Author, orderItem.Price, orderItem.ImgPath, orderItem.OrderID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func GetOrderItemsByOrderID(orderID string) ([]*model.OrderItem, error) {
	const query = "SELECT id, count, amount, title, author, price, img_path, order_id FROM order_items WHERE order_id = ?"
	rows, err := utils.DB.Query(query, orderID)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch order items: %v", err)
	}
	defer rows.Close()

	var orderItems []*model.OrderItem
	for rows.Next() {
		orderItem := &model.OrderItem{}
		if err := rows.Scan(&orderItem.OrderID, &orderItem.Count, &orderItem.Amount, &orderItem.Title, &orderItem.Author, &orderItem.Price, &orderItem.ImgPath, &orderItem.OrderID); err != nil {
			return nil, fmt.Errorf("Failed to scan order row: %v", err)
		}
		orderItems = append(orderItems, orderItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating over order items rows: %v", err)
	}
	return orderItems, nil
}
