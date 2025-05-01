package users

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserController interface {
	CreateUser(ctx *fasthttp.RequestCtx)
	GetUserByID(ctx *fasthttp.RequestCtx)
	GetUserByUsername(ctx *fasthttp.RequestCtx)
	UpdateUser(ctx *fasthttp.RequestCtx)
	DeleteUser(ctx *fasthttp.RequestCtx)
}

type UserControllerImpl struct {
	service UserService
}

func NewUserController(service UserService) *UserControllerImpl {
	return &UserControllerImpl{service: service}
}

func (c *UserControllerImpl) CreateUser(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - CreateUser endpoint called")

	var user User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("ERROR: Failed to unmarshal request body: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}

	if err := c.service.CreateUser(&user); err != nil {
		log.Printf("ERROR: Failed to create user: %v", err)
		if err == ErrUsernameExists {
			ctx.SetStatusCode(fasthttp.StatusConflict)
			ctx.SetBodyString(`{"error": "Username already exists"}`)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to create user"}`)
		return
	}

	log.Printf("DEBUG: User created successfully with ID: %v", user.ID)
	ctx.SetStatusCode(fasthttp.StatusCreated)
	json.NewEncoder(ctx).Encode(user)
}

func (c *UserControllerImpl) GetUserByID(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - GetUserByID endpoint called")

	idStr := ctx.UserValue("id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("ERROR: Invalid ID format: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid ID format"}`)
		return
	}

	user, err := c.service.GetUserByID(id)
	if err != nil {
		log.Printf("ERROR: Failed to get user: %v", err)
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString(`{"error": "User not found"}`)
		return
	}

	log.Printf("DEBUG: User found with ID: %v", id)
	json.NewEncoder(ctx).Encode(user)
}

func (c *UserControllerImpl) GetUserByUsername(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - GetUserByUsername endpoint called")

	username := ctx.UserValue("username").(string)
	user, err := c.service.GetUserByUsername(username)
	if err != nil {
		log.Printf("ERROR: Failed to get user by username: %v", err)
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString(`{"error": "User not found"}`)
		return
	}

	log.Printf("DEBUG: User found with username: %s", username)
	json.NewEncoder(ctx).Encode(user)
}

func (c *UserControllerImpl) UpdateUser(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - UpdateUser endpoint called")

	var user User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("ERROR: Failed to unmarshal request body: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}

	if err := c.service.UpdateUser(&user); err != nil {
		log.Printf("ERROR: Failed to update user: %v", err)
		if err == ErrUsernameExists {
			ctx.SetStatusCode(fasthttp.StatusConflict)
			ctx.SetBodyString(`{"error": "Username already exists"}`)
			return
		}
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to update user"}`)
		return
	}

	log.Printf("DEBUG: User updated successfully with ID: %v", user.ID)
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(user)
}

func (c *UserControllerImpl) DeleteUser(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: Controller layer - DeleteUser endpoint called")

	idStr := ctx.UserValue("id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		log.Printf("ERROR: Invalid ID format: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid ID format"}`)
		return
	}

	if err := c.service.DeleteUser(id); err != nil {
		log.Printf("ERROR: Failed to delete user: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to delete user"}`)
		return
	}

	log.Printf("DEBUG: User deleted successfully with ID: %v", id)
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(`{"message": "User deleted successfully"}`)
}
