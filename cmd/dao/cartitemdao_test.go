package dao

import (
	"fmt"
	"testing"

	"bookstore/cmd/model"
)

func TestCartItem(t *testing.T) {
	t.Run("Get cart item by book ID and cart ID validation", testGetCartItemByBookIDAndCartID)
	t.Run("Get cart items by cart ID validation", testGetCartItemsByCartID)
	t.Run("Update cart item quantity validation", testUpdateCartItemQuantity)
	t.Run("Delete cart items by cart ID validation", testDeleteCartItemsByCartID)
	t.Run("Delete cart item by ID validation", testDeleteCartItemByID)
}

func testGetCartItemByBookIDAndCartID(t *testing.T) {
	bookID := "1"
	cartID := "test"
	cartItem, err := GetCartItemByBookIDAndCartID(bookID, cartID)
	if err != nil {
		t.Fatalf("GetCartItemByBookIDAndCartID failed:%v", err)
	}
	fmt.Printf("Information for cart item: ID: %d, Count: %d, Amount: %.2f, CartID: %s\n", cartItem.ID, cartItem.Count, cartItem.Amount, cartItem.CartID)
}

func testGetCartItemsByCartID(t *testing.T) {
	cartID := "test"
	cartItems, err := GetCartItemsByCartID(cartID)
	if err != nil {
		t.Fatalf("GetCartItemsByCartID failed:%v", err)
	}

	for i, v := range cartItems {
		fmt.Printf("Information for cart item No#%d: ID: %d, Count: %d, Amount: %.2f, CartID: %s\n", i+1, v.ID, v.Count, v.Amount, v.CartID)
	}
}

func testUpdateCartItemQuantity(t *testing.T) {
	book := &model.Book{
		ID:    2,
		Title: "88首自選",
		Price: 380.00,
	}
	cartItem := &model.CartItem{
		ID:     2,
		Count:  4,
		Book:   book,
		CartID: "test",
	}

	if err := UpdateCartItemQuantity(cartItem); err != nil {
		t.Fatalf("UpdateCartItemQuantity failed: %v", err)
	}
}

func testDeleteCartItemsByCartID(t *testing.T) {
	cartID := "0507c7fe-6be0-46e9-ac87-48198f22c41c"

	if err := DeleteCartItemsByCartID(cartID); err != nil {
		t.Fatalf("DeleteCartItemsByCartID failed: %v", err)
	}
}

func testDeleteCartItemByID(t *testing.T) {
	cartItemID := "15"

	if err := DeleteCartItemByID(cartItemID); err != nil {
		t.Fatalf("DeleteCartItemsByCartID failed: %v", err)
	}
}
