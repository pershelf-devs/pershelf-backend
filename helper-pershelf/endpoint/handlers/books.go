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
	if books := crud.GetAllBooks(); books == nil {
		log.Printf("(Error): error retrieving books list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): Books list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
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
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	book = crud.GetBookByID(bookID)
	if book.ID == 0 {
		log.Printf("(Error): book not found by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): book retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(book); err != nil {
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
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	book = crud.CreateBook(&book)
	if book.ID == 0 {
		log.Printf("(Error): error creating book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): book created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
	if err := json.NewEncoder(ctx).Encode(book); err != nil {
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
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	book = crud.UpdateBook(book)
	if book.ID == 0 {
		log.Printf("(Error): error updating book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): book updated successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// DeleteBookHandler deletes a book entry from the database.
func DeleteBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteBook(bookID); err != nil {
		log.Printf("(Error): error deleting book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): book deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
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
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): book retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(book); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}
