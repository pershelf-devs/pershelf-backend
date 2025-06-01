package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllUserShelfsHandler retrieves all user_shelf entries from the database.
func GetAllUserShelfsHandler(ctx *fasthttp.RequestCtx) {
	var userShelfs []crud.UserShelf
	if userShelfs = crud.GetAllUserShelfs(); userShelfs == nil {
		log.Printf("(Error): error retrieving user_shelfs list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): user_shelfs list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
		Status:     response.ResponseMessage{Code: "0", Values: nil},
		UserShelfs: userShelfs,
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
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"Error converting user_shelf ID to int"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	userShelf = crud.GetUserShelfByID(userShelfID)
	if userShelf.ID == 0 {
		log.Printf("(Error): user_shelf not found by ID at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"User shelf not found"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): user_shelf retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
		Status:     response.ResponseMessage{Code: "0", Values: nil},
		UserShelfs: []crud.UserShelf{userShelf},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserShelvesByUserIDHandler retrieves a user_shelf entry by user ID from the database.
func GetUserShelvesByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
		userShelves []crud.UserShelf
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	userShelves = crud.GetUserShelvesByUserID(userID)
	if userShelves == nil {
		log.Printf("(Error): user_shelf not found by user ID at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"User shelf not found"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): user_shelf retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
		Status:     response.ResponseMessage{Code: "0", Values: nil},
		UserShelfs: userShelves,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserShelfByUserIDAndNameHandler retrieves a user_shelf entry by user ID and name from the database.
func GetUserShelfByUserIDAndNameHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
		name        = string(ctx.UserValue("name").(string))
		userShelf   crud.UserShelf
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userID == 0 {
		log.Printf("(Error): user ID is 0 at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"User ID is 0"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
	}

	if name == "" {
		log.Printf("(Error): name is empty at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
			Status:     response.ResponseMessage{Code: "3", Values: []string{"Name is empty"}},
			UserShelfs: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
	}

	userShelf = crud.GetUserShelfByUserIDAndName(userID, name)
	log.Printf("(Information): user_shelf retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UserShelfsResp{
		Status:     response.ResponseMessage{Code: "0", Values: nil},
		UserShelfs: []crud.UserShelf{userShelf},
	}); err != nil {
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
