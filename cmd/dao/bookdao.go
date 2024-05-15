package dao

import (
	"database/sql"
	"fmt"
	"strconv"

	"bookstore/cmd/model"
	"bookstore/cmd/utils"
)

func GetBooks() ([]*model.Book, error) {
	const query = "SELECT id, title, author, price, sales, stock, img_path FROM books"

	rows, err := utils.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch books: %v", err)
	}
	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath); err != nil {
			return nil, fmt.Errorf("Failed to scan book row: %v", err)
		}
		books = append(books, book)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error while iterating over books rows: %v", err)
	}
	return books, nil
}

func AddBook(book *model.Book) error {
	const query = "INSERT INTO books (title, author, price, sales, stock, img_path) VALUES (?,?,?,?,?,?)"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ImgPath); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func DeleteBook(bookID string) error {
	const query = "DELETE FROM books WHERE id = ?"

	if _, err := utils.DB.Exec(query, bookID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	return nil
}

func GetBookByID(bookID string) (*model.Book, error) {
	const query = "SELECT id, title, author, price, sales, stock, img_path FROM books WHERE id = ?"
	row := utils.DB.QueryRow(query, bookID)

	book := &model.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath)
	switch {
	case err == sql.ErrNoRows:
		return nil, fmt.Errorf("Book not found")
	case err != nil:
		return nil, fmt.Errorf("Failed to scan book row: %v", err)
	}

	return book, nil
}

func UpdateBook(book *model.Book) error {
	const query = "UPDATE books SET title=?, author=?, price=?, sales=?, stock=? WHERE id=?"
	stmt, err := utils.DB.Prepare(query)
	if err != nil {
		return fmt.Errorf("Failed to prepare statement: %v", err)
	}

	if _, err := stmt.Exec(book.Title, book.Author, book.Price, book.Sales, book.Stock, book.ID); err != nil {
		return fmt.Errorf("Failed to execute statement: %v", err)
	}
	defer stmt.Close()
	return nil
}

func GetPaginatedBooks(pageNo string) (*model.PaginatedBook, error) {
	const queryCount = "SELECT COUNT(*) FROM books"
	intPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	var totalRecordCount int64

	row := utils.DB.QueryRow(queryCount)
	if err := row.Scan(&totalRecordCount); err != nil {
		return nil, fmt.Errorf("Failed to scan total record count: %v", err)
	}

	const pageSize = 4
	totalPageNo := totalRecordCount / pageSize
	if totalRecordCount%pageSize != 0 {
		totalPageNo++
	}

	const queryBooks = "SELECT id, title, author, price, sales, stock, img_path FROM books LIMIT ?,?"
	stmt, err := utils.DB.Prepare(queryBooks)
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query((intPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch books: %v", err)
	}
	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath); err != nil {
			return nil, fmt.Errorf("Failed to scan book row: %v", err)
		}
		books = append(books, book)
	}

	paginatedBooks := &model.PaginatedBook{
		Books:            books,
		PageNo:           intPageNo,
		PageSize:         pageSize,
		TotalPageNo:      totalPageNo,
		TotalRecordCount: totalRecordCount,
	}
	return paginatedBooks, nil
}

func FetchBooksByPriceRange(pageNo, minPrice, maxPrice string) (*model.PaginatedBook, error) {
	const queryCount = "SELECT COUNT(*) FROM books WHERE price BETWEEN ? AND ?"
	intPageNo, _ := strconv.ParseInt(pageNo, 10, 64)
	var totalRecordCount int64

	row := utils.DB.QueryRow(queryCount, minPrice, maxPrice)
	if err := row.Scan(&totalRecordCount); err != nil {
		return nil, fmt.Errorf("Failed to scan total record count: %v", err)
	}

	const pageSize = 4
	totalPageNo := totalRecordCount / pageSize
	if totalRecordCount%pageSize != 0 {
		totalPageNo++
	}

	const queryBooks = "SELECT id, title, author, price, sales, stock, img_path FROM books WHERE price BETWEEN ? AND ? LIMIT ?, ?"
	stmt, err := utils.DB.Prepare(queryBooks)
	if err != nil {
		return nil, fmt.Errorf("Failed to prepare statement: %v", err)
	}
	defer stmt.Close()

	rows, err := stmt.Query(minPrice, maxPrice, (intPageNo-1)*pageSize, pageSize)
	if err != nil {
		return nil, fmt.Errorf("Failed to fetch books: %v", err)
	}
	defer rows.Close()

	var books []*model.Book
	for rows.Next() {
		book := &model.Book{}
		if err := rows.Scan(&book.ID, &book.Title, &book.Author, &book.Price, &book.Sales, &book.Stock, &book.ImgPath); err != nil {
			return nil, fmt.Errorf("Failed to scan book row: %v", err)
		}
		books = append(books, book)
	}

	paginatedBooks := &model.PaginatedBook{
		Books:            books,
		PageNo:           intPageNo,
		PageSize:         pageSize,
		TotalPageNo:      totalPageNo,
		TotalRecordCount: totalRecordCount,
	}
	return paginatedBooks, nil
}
