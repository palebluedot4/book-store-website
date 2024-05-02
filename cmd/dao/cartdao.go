package dao

import (
	"bookstore/cmd/model"
	"bookstore/cmd/utils"
	"database/sql"
	"fmt"
)

func AddCart(cart *model.Cart) error {
	const query = "INSERT INTO carts (id, total_count, total_amount, user_id) VALUES (?, ?, ?, ?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(cart.CartID, cart.GetTotalCount(), cart.GetTotalAmount(), cart.UserID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()

	for _, v := range cart.CartItems {
		if err := AddCartItem(v); err != nil {
			return fmt.Errorf("Failed to add cart item: %v", err)
		}
	}
	return nil
}

func GetCartByUserID(userID int64) (*model.Cart, error) {
	const query = "SELECT id, total_count, total_amount, user_id FROM carts WHERE user_id = ?"
	row := utils.DB.QueryRow(query, userID)

	cart := &model.Cart{}
	err := row.Scan(&cart.CartID, &cart.TotalCount, &cart.TotalAmount, &cart.UserID)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("Cart not found")
	case err != nil:
		return nil, fmt.Errorf("Failed to scan cart row: %v", err)
	}

	cart.CartItems, _ = GetCartItemsByCartID(cart.CartID)
	return cart, nil
}

func UpdateCart(cart *model.Cart) error {
	const query = "UPDATE carts SET total_count = ?, total_amount = ? WHERE id = ?"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(cart.GetTotalCount(), cart.GetTotalAmount(), cart.CartID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func DeleteCartByCartID(cartID string) error {
	if err := DeleteCartItemsByCartID(cartID); err != nil {
		return fmt.Errorf("Failed to delete cart items: %v", err)
	}

	const query = "DELETE FROM carts WHERE id = ?"
	if _, err := utils.DB.Exec(query, cartID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	return nil
}
