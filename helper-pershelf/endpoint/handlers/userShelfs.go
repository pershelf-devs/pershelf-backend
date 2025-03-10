package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllUserShelvesHandler retrieves all user_shelf entries from the database.
func GetAllUserShelvesHandler(ctx *fasthttp.RequestCtx) {
	var userShelves []crud.UserShelf
	if userShelves = crud.GetAllUserShelves(); userShelves == nil {
		log.Printf("(Error): error retrieving user_shelves list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user_shelves list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(response.UserShelvesResp{
		Status:      response.ResponseMessage{Code: "0", Values: nil},
		UserShelves: userShelves,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetUserShelfByIDHandler retrieves a user_shelf entry by ID from the database.
func GetUserShelfByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth              = ctx.Path()
		userShelfID, err = strconv.Atoi(ctx.UserValue("id").(string))
		userShelf        crud.UserShelf
	)

	if err != nil {
		log.Printf("(Error): error converting user_shelf ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userShelf = crud.GetUserShelfByID(userShelfID)
	if userShelf.ID == 0 {
		log.Printf("(Error): user_shelf not found by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): user_shelf retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(userShelf); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateUserShelfHandler creates a new user_shelf entry in the database.
func CreateUserShelfHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth       = ctx.Path()
		userShelf crud.UserShelf
	)

	if err := json.Unmarshal(ctx.Request.Body(), &userShelf); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userShelf = crud.CreateUserShelf(&userShelf)
	if userShelf.ID == 0 {
		log.Printf("(Error): error creating user_shelf at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user_shelf created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// UpdateUserShelfHandler updates an existing user_shelf entry in the database.
func UpdateUserShelfHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth       = ctx.Path()
		userShelf crud.UserShelf
	)

	if err := json.Unmarshal(ctx.Request.Body(), &userShelf); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	userShelf = crud.UpdateUserShelf(userShelf)
	if userShelf.ID == 0 {
		log.Printf("(Error): error updating user_shelf at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user_shelf updated successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// DeleteUserShelfHandler deletes a user_shelf entry from the database.
func DeleteUserShelfHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth              = ctx.Path()
		userShelfID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting user_shelf ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteUserShelf(userShelfID); err != nil {
		log.Printf("(Error): error deleting user_shelf at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user_shelf deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
