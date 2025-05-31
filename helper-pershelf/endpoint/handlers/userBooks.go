package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllUserBooksHandler retrieves all user books from the database and sends them in a response to the client's request.
func GetAllUserBooksHandler(ctx *fasthttp.RequestCtx) {
	var userBooks []crud.UserBook
	if userBooks = crud.GetAllUserBooks(); userBooks == nil {
		log.Printf("(Error): error retrieving user books list at endpoint (%s).", string(ctx.Path()))
		if err := json.NewEncoder(ctx).Encode(response.UserBooksResp{
			Status:    response.ResponseMessage{Code: "4", Values: nil},
			UserBooks: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
		}
		return
	}

	log.Printf("(Information): user books list retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserBooksResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		UserBooks: userBooks,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetUserBookByIDHandler retrieves a user book by ID from the database.
func GetUserBookByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("id").(string))
		userBook    crud.UserBook
	)

	if err != nil {
		log.Printf("(Error): error converting user book id to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userBook = crud.GetUserBookByID(bookID)
	if userBook.ID == 0 {
		log.Printf("(Error): error retrieving user book by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): user book retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserBooksResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		UserBooks: []crud.UserBook{userBook},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserBooksByUserIDHandler retrieves all books associated with a user ID.
func GetUserBooksByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
		userBooks   []crud.UserBook
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userBooks = crud.GetUserBooksByUserID(userID)
	if userBooks == nil {
		log.Printf("(Error): error retrieving user books by user ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): user books retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserBooksResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		UserBooks: userBooks,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserBooksByBookIDHandler retrieves all users associated with a specific book ID.
func GetUserBooksByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("book-id").(string))
		userBooks   []crud.UserBook
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userBooks = crud.GetUserBookByBookID(bookID)
	if userBooks == nil {
		log.Printf("(Error): error retrieving user books by book ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): user books retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserBooksResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		UserBooks: userBooks,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateUserBookHandler creates a new user book entry in the database.
func CreateUserBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth      = ctx.Path()
		userBook crud.UserBook
	)

	if err := json.Unmarshal(ctx.Request.Body(), &userBook); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userBook = crud.CreateUserBook(&userBook)
	if userBook.ID == 0 {
		log.Printf("(Error): error creating user book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user book created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// UpdateUserBookHandler updates a user book entry in the database.
func UpdateUserBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth      = ctx.Path()
		userBook crud.UserBook
	)

	if err := json.Unmarshal(ctx.Request.Body(), &userBook); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userBook = crud.UpdateUserBook(&userBook)
	if userBook.ID == 0 {
		log.Printf("(Error): error updating user book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user book updated successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// DeleteUserBookHandler deletes a user book entry from the database.
func DeleteUserBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting user book ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteUserBook(bookID); err != nil {
		log.Printf("(Error): error deleting user book at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user book deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
