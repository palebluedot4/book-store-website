package dao

import (
	"bookstore/cmd/model"
	"bookstore/cmd/utils"
	"database/sql"
	"fmt"
)

func AddCartItem(cartItem *model.CartItem) error {
	const query = "INSERT INTO cart_items (count, amount, book_id, cart_id) VALUES (?, ?, ?, ?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func GetCartItemByBookIDAndCartID(bookID, cartID string) (*model.CartItem, error) {
	const query = "SELECT id, count, amount, cart_id FROM cart_items WHERE book_id = ? AND cart_id = ?"
	row := utils.DB.QueryRow(query, bookID, cartID)

	cartItem := &model.CartItem{}
	err := row.Scan(&cartItem.ID, &cartItem.Count, &cartItem.Amount, &cartItem.CartID)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("Cart item not found")
	case err != nil:
		return nil, fmt.Errorf("Failed to scan cart item row: %v", err)
	}

	book, _ := GetBookByID(bookID)
	cartItem.Book = book
	return cartItem, nil
}

func GetCartItemsByCartID(cartID string) ([]*model.CartItem, error) {
	const query = "SELECT id, count, amount, book_id, cart_id FROM cart_items WHERE cart_id = ?"

	rows, err := utils.DB.Query(query, cartID)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch cart items: %v", err)
	}
	defer rows.Close()

	var cartItems []*model.CartItem
	for rows.Next() {
		var bookID string
		cartItem := &model.CartItem{}
		if err := rows.Scan(&cartItem.ID, &cartItem.Count, &cartItem.Amount, &bookID, &cartItem.CartID); err != nil {
			return nil, fmt.Errorf("Failed to scan cart items row: %v", err)
		}
		book, _ := GetBookByID(bookID)
		cartItem.Book = book
		cartItems = append(cartItems, cartItem)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating over cart items rows: %v", err)
	}
	return cartItems, nil
}

func UpdateCartItemQuantity(cartItem *model.CartItem) error {
	const query = "UPDATE cart_items SET count = ?, amount = ? WHERE book_id = ? AND cart_id = ?"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(cartItem.Count, cartItem.GetAmount(), cartItem.Book.ID, cartItem.CartID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func DeleteCartItemsByCartID(cartID string) error {
	const query = "DELETE FROM cart_items WHERE cart_id = ?"

	if _, err := utils.DB.Exec(query, cartID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	return nil
}

func DeleteCartItemByID(cartItemID string) error {
	const query = "DELETE FROM cart_items WHERE id = ?"

	if _, err := utils.DB.Exec(query, cartItemID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	return nil
}
