package dao

import (
	"bookstore/cmd/model"
	"fmt"
	"testing"
)

func TestBook(t *testing.T) {
	// t.Run("Get books validation", testGetBooks)
	// t.Run("Add book validation", testAddBook)
	// t.Run("Delete book validation", testDeleteBook)
	// t.Run("Get book by ID validation", testGetBookByID)
	// t.Run("Update book validation", testUpdateBook)
	// t.Run("Get paginated books validation", testGetPaginatedBooks)
	// t.Run("Fetch books by price range validation", testFetchBooksByPriceRange)
}

func testGetBooks(t *testing.T) {
	books, err := GetBooks()
	if err != nil {
		t.Fatalf("GetBooks failed: %v", err)
	}

	for i, v := range books {
		fmt.Printf("Information for book #%d: %v\n", i+1, v)
	}
}

func testAddBook(t *testing.T) {
	book := &model.Book{
		Title:   "深入淺出Python",
		Author:  "Paul Barry",
		Price:   880.00,
		Sales:   100,
		Stock:   100,
		ImgPath: "static/img/default.jpg",
	}

	if err := AddBook(book); err != nil {
		t.Fatalf("AddBook failed: %v", err)
	}
}

func testDeleteBook(t *testing.T) {
	bookID := "35"

	if err := DeleteBook(bookID); err != nil {
		t.Fatalf("DeleteBook failed: %v", err)
	}
}

func testGetBookByID(t *testing.T) {
	bookID := "24"

	book, err := GetBookByID(bookID)
	if err != nil {
		t.Fatalf("GetBookByID failed: %v", err)
	}
	fmt.Printf("Information for book #%s: %v", bookID, book)
}

func testUpdateBook(t *testing.T) {
	book := &model.Book{
		ID:      23,
		Title:   "遠處的拉莫",
		Author:  "胡遷",
		Price:   360.00,
		Sales:   120,
		Stock:   80,
		ImgPath: "static/img/default.jpg",
	}

	if err := UpdateBook(book); err != nil {
		t.Fatalf("UpdateBook failed: %v", err)
	}
}

func testGetPaginatedBooks(t *testing.T) {
	PageNo := "2"
	paginatedBooks, err := GetPaginatedBooks(PageNo)
	if err != nil {
		t.Fatalf("GetPaginatedBooks failed: %v", err)
	}
	fmt.Printf("Information for page: page#%v (total pages: %v, total records: %v)\n", paginatedBooks.PageNo, paginatedBooks.TotalPageNo, paginatedBooks.TotalRecordCount)

	for _, v := range paginatedBooks.Books {
		fmt.Printf("Information for paginated books: %v\n", v)
	}
}

func testFetchBooksByPriceRange(t *testing.T) {
	PageNo := "1"
	minPrice := "400"
	maxPrice := "800"
	paginatedBooks, err := FetchBooksByPriceRange(PageNo, minPrice, maxPrice)
	if err != nil {
		t.Fatalf("FetchBooksByPriceRange failed: %v", err)
	}
	fmt.Printf("Information for page#%v: total pages: %v, total records: %v\n", paginatedBooks.PageNo, paginatedBooks.TotalPageNo, paginatedBooks.TotalRecordCount)

	for _, v := range paginatedBooks.Books {
		fmt.Printf("Information for fetched books: %v\n", v)
	}
}
