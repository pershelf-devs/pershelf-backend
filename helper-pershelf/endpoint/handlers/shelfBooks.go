package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllShelfBooksHandler retrieves all shelf_book entries from the database.
func GetAllShelfBooksHandler(ctx *fasthttp.RequestCtx) {
	var shelfBooks []crud.ShelfBook
	if shelfBooks = crud.GetAllShelfBooks(); shelfBooks == nil {
		log.Printf("(Error): error retrieving shelf_books list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): shelf_books list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(response.ShelfBooksResp{
		Status:     response.ResponseMessage{Code: "0", Values: nil},
		ShelfBooks: shelfBooks,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetShelfBookByIDHandler retrieves a shelf_book entry by ID from the database.
func GetShelfBookByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth              = ctx.Path()
		shelfBookID, err = strconv.Atoi(ctx.UserValue("id").(string))
		shelfBook        crud.ShelfBook
	)

	if err != nil {
		log.Printf("(Error): error converting shelf_book ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	shelfBook = crud.GetShelfBookByID(shelfBookID)
	if shelfBook.ID == 0 {
		log.Printf("(Error): shelf_book not found by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): shelf_book retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(shelfBook); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateShelfBookHandler creates a new shelf_book entry in the database.
func CreateShelfBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth       = ctx.Path()
		shelfBook crud.ShelfBook
	)

	if err := json.Unmarshal(ctx.Request.Body(), &shelfBook); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	shelfBook = crud.CreateShelfBook(&shelfBook)
	if shelfBook.ID == 0 {
		log.Printf("(Error): error creating shelf_book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): shelf_book created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// DeleteShelfBookHandler deletes a shelf_book entry from the database.
func DeleteShelfBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth              = ctx.Path()
		shelfBookID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting shelf_book ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteShelfBook(shelfBookID); err != nil {
		log.Printf("(Error): error deleting shelf_book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): shelf_book deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
