package dao

import (
	"fmt"

	"bookstore/cmd/model"
	"bookstore/cmd/utils"
)

func AddOrder(order *model.Order) error {
	const query = "INSERT INTO orders (id, create_time, total_count, total_amount, status, user_id) VALUES (?, ?, ?, ?, ?, ?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(order.OrderID, order.CreateTime, order.TotalCount, order.TotalAmount, order.Status, order.UserID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func GetOrders() ([]*model.Order, error) {
	const query = "SELECT id, create_time, total_count, total_amount, status, user_id FROM orders"
	rows, err := utils.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch orders: %v", err)
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		if err := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.Status, &order.UserID); err != nil {
			return nil, fmt.Errorf("Failed to scan order row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating over orders rows: %v", err)
	}
	return orders, nil
}

func GetMyOrders(userID int64) ([]*model.Order, error) {
	const query = "SELECT id, create_time, total_count, total_amount, status, user_id FROM orders WHERE user_id = ?"
	rows, err := utils.DB.Query(query, userID)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch orders: %v", err)
	}
	defer rows.Close()

	var orders []*model.Order
	for rows.Next() {
		order := &model.Order{}
		if err := rows.Scan(&order.OrderID, &order.CreateTime, &order.TotalCount, &order.TotalAmount, &order.Status, &order.UserID); err != nil {
			return nil, fmt.Errorf("Failed to scan order row: %v", err)
		}
		orders = append(orders, order)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating over orders rows: %v", err)
	}
	return orders, nil
}

func UpdateOrderStatus(orderID string, status model.OrderStatus) error {
	const query = "UPDATE orders SET status = ? WHERE id = ?"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(status, orderID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}
