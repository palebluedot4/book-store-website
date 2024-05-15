package controller

import (
	"html/template"
	"log"
	"net/http"
	"strconv"

	"bookstore/cmd/dao"
	"bookstore/cmd/model"
)

// func IndexHandler(w http.ResponseWriter, r *http.Request) {
// 	pageNo := r.FormValue("pageNo")

// 	if pageNo == "" {
// 		pageNo = "1"
// 	}

// page, err := dao.GetPaginatedBooks(pageNo)
// if err != nil {
// 	log.Printf("GetPaginatedBooks failed: %v", err)
// 	http.Error(w, "Failed to retrieve paginated books", http.StatusInternalServerError)
// 	return
// }

// tmpl := template.Must(template.ParseFiles("views/index.html"))
// if err := tmpl.Execute(w, page); err != nil {
// 	log.Printf("Failed to execute IndexHandler template: %s", err)
// 	http.Error(w, "Failed to render template", http.StatusInternalServerError)
// 	return
// }
// }

// func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
// 	books, err := dao.GetBooks()
// 	if err != nil {
// 		log.Printf("GetBooksHandler failed: %v", err)
// 		http.Error(w, "failed to retrieve books", http.StatusInternalServerError)
// 		return
// 	}

// 	tmpl := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
// 	if err := tmpl.Execute(w, books); err != nil {
// 		log.Printf("Failed to execute GetBooksHandler template: %s", err)
// 		http.Error(w, "Failed to render template", http.StatusInternalServerError)
// 		return
// 	}
// }

// func AddBookHandler(w http.ResponseWriter, r *http.Request) {
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")

// 	floatPrice, _ := strconv.ParseFloat(price, 64)
// 	intSales, _ := strconv.ParseInt(sales, 10, 0)
// 	intStock, _ := strconv.ParseInt(stock, 10, 0)

// 	book := &model.Book{
// 		Title:   title,
// 		Author:  author,
// 		Price:   floatPrice,
// 		Sales:   intSales,
// 		Stock:   intStock,
// 		ImgPath: "static/img/default.jpg",
// 	}

// 	if err := dao.AddBook(book); err != nil {
// 		log.Printf("AddBook failed: %v", err)
// 		http.Error(w, "Failed to add book", http.StatusInternalServerError)
// 		return
// 	}

// 	GetPaginatedBooksHandler(w, r)
// }

func DeleteBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")

	if err := dao.DeleteBook(bookID); err != nil {
		log.Printf("DeleteBookHandler failed: %v", err)
		http.Error(w, "Failed to delete book", http.StatusInternalServerError)
		return
	}

	GetPaginatedBooksHandler(w, r)
}

func EditBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := r.FormValue("bookID")

	book, _ := dao.GetBookByID(bookID)
	switch {
	case book != nil:
		tmpl := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		if err := tmpl.Execute(w, book); err != nil {
			log.Printf("Failed to execute EditBookHandler template: %s", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	default:
		tmpl := template.Must(template.ParseFiles("views/pages/manager/book_edit.html"))
		if err := tmpl.Execute(w, nil); err != nil {
			log.Printf("Failed to execute EditBookHandler template: %s", err)
			http.Error(w, "Failed to render template", http.StatusInternalServerError)
			return
		}
	}
}

// func UpdateBookHandler(w http.ResponseWriter, r *http.Request) {
// 	bookID := r.PostFormValue("bookID")
// 	title := r.PostFormValue("title")
// 	author := r.PostFormValue("author")
// 	price := r.PostFormValue("price")
// 	sales := r.PostFormValue("sales")
// 	stock := r.PostFormValue("stock")

// 	intBookID, _ := strconv.ParseInt(bookID, 10, 0)
// 	floatPrice, _ := strconv.ParseFloat(price, 64)
// 	intSales, _ := strconv.ParseInt(sales, 10, 0)
// 	intStock, _ := strconv.ParseInt(stock, 10, 0)

// 	book := &model.Book{
// 		ID:      intBookID,
// 		Title:   title,
// 		Author:  author,
// 		Price:   floatPrice,
// 		Sales:   intSales,
// 		Stock:   intStock,
// 		ImgPath: "static/img/default.jpg",
// 	}

// 	if err := dao.UpdateBook(book); err != nil {
// 		log.Printf("UpdateBook failed: %v", err)
// 		http.Error(w, "Failed to update book: ", http.StatusInternalServerError)
// 		return
// 	}

// 	GetPaginatedBooksHandler(w, r)
// }

func UpdateOrAddBookHandler(w http.ResponseWriter, r *http.Request) {
	bookID := r.PostFormValue("bookID")
	title := r.PostFormValue("title")
	author := r.PostFormValue("author")
	price := r.PostFormValue("price")
	sales := r.PostFormValue("sales")
	stock := r.PostFormValue("stock")

	intBookID, _ := strconv.ParseInt(bookID, 10, 64)
	floatPrice, _ := strconv.ParseFloat(price, 64)
	intSales, _ := strconv.ParseInt(sales, 10, 64)
	intStock, _ := strconv.ParseInt(stock, 10, 64)

	book := &model.Book{
		ID:      intBookID,
		Title:   title,
		Author:  author,
		Price:   floatPrice,
		Sales:   intSales,
		Stock:   intStock,
		ImgPath: "static/img/default.jpg",
	}

	switch {
	case book.ID > 0:
		if err := dao.UpdateBook(book); err != nil {
			log.Printf("UpdateBook failed: %v", err)
			http.Error(w, "Failed to update or add book: ", http.StatusInternalServerError)
			return
		}
	default:
		if err := dao.AddBook(book); err != nil {
			log.Printf("AddBook failed: %v", err)
			http.Error(w, "Failed to add book", http.StatusInternalServerError)
			return
		}
	}

	GetPaginatedBooksHandler(w, r)
}

func GetPaginatedBooksHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")

	if pageNo == "" {
		pageNo = "1"
	}

	page, err := dao.GetPaginatedBooks(pageNo)
	if err != nil {
		log.Printf("GetPaginatedBooks failed: %v", err)
		http.Error(w, "Failed to retrieve paginated books", http.StatusInternalServerError)
		return
	}

	tmpl := template.Must(template.ParseFiles("views/pages/manager/book_manager.html"))
	if err := tmpl.Execute(w, page); err != nil {
		log.Printf("Failed to execute GetPaginatedBooksHandler template: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}

func FetchBooksByPriceRangeHandler(w http.ResponseWriter, r *http.Request) {
	pageNo := r.FormValue("pageNo")
	minPrice := r.FormValue("min")
	maxPrice := r.FormValue("max")
	var page *model.PaginatedBook
	var err error

	if pageNo == "" {
		pageNo = "1"
	}
	switch {
	case minPrice == "" && maxPrice == "":
		page, err = dao.GetPaginatedBooks(pageNo)
		if err != nil {
			log.Printf("GetPaginatedBooks failed: %v", err)
			http.Error(w, "Failed to retrieve paginated books", http.StatusInternalServerError)
			return
		}
	default:
		page, err = dao.FetchBooksByPriceRange(pageNo, minPrice, maxPrice)
		if err != nil {
			log.Printf("FetchBooksByPriceRange failed: %v", err)
			http.Error(w, "Failed to fetch books by price range", http.StatusInternalServerError)
			return
		}
		page.MinPrice = minPrice
		page.MaxPrice = maxPrice
	}

	loggedIn, session := dao.IsLoggedIn(r)
	if loggedIn {
		page.IsLoggedIn = true
		page.Username = session.Username
	}

	tmpl := template.Must(template.ParseFiles("views/index.html"))
	if err := tmpl.Execute(w, page); err != nil {
		log.Printf("Failed to execute FetchBooksByPriceRangeHandler template: %s", err)
		http.Error(w, "Failed to render template", http.StatusInternalServerError)
		return
	}
}
