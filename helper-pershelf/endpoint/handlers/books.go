package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllBooksHandler retrieves all books from the database.
func GetAllBooksHandler(ctx *fasthttp.RequestCtx) {
	var books []crud.Book
	if books = crud.GetAllBooks(); books == nil {
		log.Printf("(Error): error retrieving books list at endpoint (%s).", string(ctx.Path()))
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Error retrieving books list"}},
			Books:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
		}
		return
	}

	log.Printf("(Information): Books list retrieved successfully (length: %d).", len(books))
	if err := json.NewEncoder(ctx).Encode(response.BooksResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Books:  books,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetBookByIDHandler retrieves a book by ID from the database.
func GetBookByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("id").(string))
		book        crud.Book
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}},
			Books:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	book = crud.GetBookByID(bookID)
	if book.ID == 0 {
		log.Printf("(Error): book not found by ID at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Book not found"}},
			Books:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BooksResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Books:  []crud.Book{book},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateBookHandler creates a new book entry in the database.
func CreateBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = ctx.Path()
		book crud.Book
	)

	if err := json.Unmarshal(ctx.Request.Body(), &book); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			Books:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.CreateBook(&book); err != nil {
		log.Printf("(Error): error creating book at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Error creating book"}},
			Books:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book created successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BooksResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Books:  []crud.Book{book},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// UpdateBookHandler updates an existing book entry in the database.
func UpdateBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = ctx.Path()
		book crud.Book
	)

	if err := json.Unmarshal(ctx.Request.Body(), &book); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	book = crud.UpdateBook(book)
	if book.ID == 0 {
		log.Printf("(Error): error updating book at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error updating book"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book updated successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteBookHandler deletes a book entry from the database.
func DeleteBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteBook(bookID); err != nil {
		log.Printf("(Error): error deleting book at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting book"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book deleted successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetBooksByISBNHandler retrieves a book by ISBN from the database.
func GetBookByISBNHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = ctx.Path()
		isbn = ctx.UserValue("isbn").(string)
		book crud.Book
	)

	book = crud.GetBookByISBN(isbn)
	if book.ID == 0 {
		log.Printf("(Error): book not found with ISBN (%s) at endpoint (%s).", isbn, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Book not found"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BooksResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Books:  []crud.Book{book},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

func GetBooksByGenreHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = ctx.Path()
		isbn = ctx.UserValue("genre").(string)
	)

	bookList := crud.GetBooksByGenre(isbn)
	if len(bookList) == 0 {
		log.Printf("(Error): no books found with genre (%s) at endpoint (%s).", isbn, string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"No books found"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): books retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BooksResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Books:  bookList,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}
