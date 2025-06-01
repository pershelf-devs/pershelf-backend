package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllUsersHandler retrieves all users from the database and sends them in a response to client's request.
func GetAllUsersHandler(ctx *fasthttp.RequestCtx) {
	var users []crud.User
	if users = crud.GetAllUsers(); users == nil {
		log.Printf("(Error): error retrieving users list at endpoint (%s).", string(ctx.Path()))
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "4", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
		}
		return
	}

	log.Printf("(Information): users list retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UsersResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Users:  users,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetUserByIDHandler retrieves a user by id from the database and sends them in a response to client's request.
func GetUserByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("id").(string))
		User        crud.User
	)

	if err != nil {
		log.Printf("(Error): error converting user id to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "4", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	User = crud.GetUserByID(userID)
	if User.ID == 0 {
		log.Printf("(Error): error retrieving user by id at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "3", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): user retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UsersResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Users:  []crud.User{User},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetUserByEmailHandler retrieves a user by email from the database and sends them in a response to client's request.
func GetUserByEmailHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth   = ctx.Path()
		email = ""
		user  crud.User
	)

	if err := json.Unmarshal(ctx.Request.Body(), &email); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "3", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	if email == "" {
		log.Printf("(Error): error retrieving user by email at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "3", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	if user = crud.GetUserByEmail(email); user.ID == 0 {
		log.Printf("(Error): error retrieving user by email at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "3", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	log.Printf("(Information): user retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.UsersResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Users:  []crud.User{user},
	}); err != nil {
		log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
	}
}

// CreateUserHandler creates a new user in the database and sends them in a response to client's request.
func CreateUserHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = ctx.Path()
		user crud.User
	)

	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{Status: response.ResponseMessage{Code: "3", Values: nil}}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	user = crud.CreateUser(&user)
	if user.ID == 0 {
		log.Printf("(Error): error creating user at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "4", Values: nil},
			Users:  nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): user created successfully. %s", user.Name)
	if err := json.NewEncoder(ctx).Encode(response.UsersResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Users:  []crud.User{user},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// UpdateUserHandler updates a user in the database and sends them in a response to client's request.
func UpdateUserHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = ctx.Path()
		user crud.User
	)

	if err := json.Unmarshal(ctx.Request.Body(), &user); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	user = crud.UpdateUser(user)
	if user.ID == 0 {
		log.Printf("(Error): error updating user at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Values: nil}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): user updated successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteUserHandler deletes a user from the database and sends them in a response to client's request.
func DeleteUserHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting user id to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	if userID == 0 {
		log.Printf("(Error): error retrieving user by id at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	if err := crud.DeleteUser(userID); err != nil {
		log.Printf("(Error): error deleting user at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4", Values: nil}); err != nil {
			log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
		}
		return
	}

	log.Printf("(Information): user deleted successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("Error encoding response for endpoint (%s): %v", pth, err)
	}
}
