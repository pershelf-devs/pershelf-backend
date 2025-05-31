package user_books

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserBookController interface {
	CreateUserBook(ctx *fasthttp.RequestCtx)
	GetUserBookByID(ctx *fasthttp.RequestCtx)
	GetUserBooksByUserID(ctx *fasthttp.RequestCtx)
	UpdateUserBook(ctx *fasthttp.RequestCtx)
	DeleteUserBook(ctx *fasthttp.RequestCtx)
}

type UserBookControllerImpl struct {
	service UserBookService
}

func NewUserBookController(service UserBookService) *UserBookControllerImpl {
	return &UserBookControllerImpl{service: service}
}

func (c *UserBookControllerImpl) CreateUserBook(ctx *fasthttp.RequestCtx) {
	var userBook UserBook
	if err := json.Unmarshal(ctx.PostBody(), &userBook); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}
	if err := c.service.CreateUserBook(&userBook); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to create user_book"}`)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusCreated)
	json.NewEncoder(ctx).Encode(userBook)
}

func (c *UserBookControllerImpl) GetUserBookByID(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid ID format"}`)
		return
	}
	userBook, err := c.service.GetUserBookByID(id)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		ctx.SetBodyString(`{"error": "UserBook not found"}`)
		return
	}
	json.NewEncoder(ctx).Encode(userBook)
}

func (c *UserBookControllerImpl) GetUserBooksByUserID(ctx *fasthttp.RequestCtx) {
	userIDStr := ctx.UserValue("user_id").(string)
	userID, err := primitive.ObjectIDFromHex(userIDStr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid user ID format"}`)
		return
	}
	userBooks, err := c.service.GetUserBooksByUserID(userID)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to get user_books"}`)
		return
	}
	json.NewEncoder(ctx).Encode(userBooks)
}

func (c *UserBookControllerImpl) UpdateUserBook(ctx *fasthttp.RequestCtx) {
	var userBook UserBook
	if err := json.Unmarshal(ctx.PostBody(), &userBook); err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}
	if err := c.service.UpdateUserBook(&userBook); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to update user_book"}`)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(userBook)
}

func (c *UserBookControllerImpl) DeleteUserBook(ctx *fasthttp.RequestCtx) {
	idStr := ctx.UserValue("id").(string)
	id, err := primitive.ObjectIDFromHex(idStr)
	if err != nil {
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid ID format"}`)
		return
	}
	if err := c.service.DeleteUserBook(id); err != nil {
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to delete user_book"}`)
		return
	}
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetBodyString(`{"message": "UserBook deleted successfully"}`)
}
