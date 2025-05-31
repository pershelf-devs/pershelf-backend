package tablesModels

import (
	"context"
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	CreateUser(user *User) error
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id primitive.ObjectID) error
}

type UserService interface {
	CreateUser(user *User) error
	GetUserByID(id primitive.ObjectID) (*User, error)
	GetUserByUsername(username string) (*User, error)
	UpdateUser(user *User) error
	DeleteUser(id primitive.ObjectID) error
}

type UserController interface {
	CreateUser(ctx *fasthttp.RequestCtx)
	GetUserByID(ctx *fasthttp.RequestCtx)
	GetUserByUsername(ctx *fasthttp.RequestCtx)
	UpdateUser(ctx *fasthttp.RequestCtx)
	DeleteUser(ctx *fasthttp.RequestCtx)
}

type UserRouter interface {
	RegisterRoutes(router *fasthttp.RequestHandler)
}

type UserRepositoryMongo struct {
	collection *mongo.Collection
}

func NewUserRepositoryMongo(collection *mongo.Collection) *UserRepositoryMongo {
	return &UserRepositoryMongo{collection: collection}
}

func (r *UserRepositoryMongo) CreateUser(user *User) error {
	_, err := r.collection.InsertOne(context.Background(), user)
	return err
}

func (r *UserRepositoryMongo) GetUserByID(id primitive.ObjectID) (*User, error) {
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"_id": id}).Decode(&user)
	return &user, err
}

func (r *UserRepositoryMongo) GetUserByUsername(username string) (*User, error) {
	var user User
	err := r.collection.FindOne(context.Background(), bson.M{"username": username}).Decode(&user)
	return &user, err
}

func (r *UserRepositoryMongo) UpdateUser(user *User) error {
	_, err := r.collection.UpdateOne(context.Background(), bson.M{"_id": user.ID}, bson.M{"$set": user})
	return err
}

func (r *UserRepositoryMongo) DeleteUser(id primitive.ObjectID) error {
	_, err := r.collection.DeleteOne(context.Background(), bson.M{"_id": id})
	return err
}

type UserControllerImpl struct {
	service UserService
}

func NewUserController(service UserService) *UserControllerImpl {
	return &UserControllerImpl{service: service}
}

func (c *UserControllerImpl) CreateUser(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: CreateUser endpoint called")

	var user User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("ERROR: Failed to unmarshal request body: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}

	if err := c.service.CreateUser(&user); err != nil {
		log.Printf("ERROR: Failed to create user: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to create user"}`)
		return
	}

	log.Printf("DEBUG: User created successfully with ID: %v", user.ID)
	ctx.SetStatusCode(fasthttp.StatusCreated)
	json.NewEncoder(ctx).Encode(user)
}

func (c *UserControllerImpl) GetUserByID(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: GetUserByID endpoint called")

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
	log.Printf("DEBUG: GetUserByUsername endpoint called")

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
	log.Printf("DEBUG: UpdateUser endpoint called")

	var user User
	if err := json.Unmarshal(ctx.PostBody(), &user); err != nil {
		log.Printf("ERROR: Failed to unmarshal request body: %v", err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		ctx.SetBodyString(`{"error": "Invalid request body"}`)
		return
	}

	if err := c.service.UpdateUser(&user); err != nil {
		log.Printf("ERROR: Failed to update user: %v", err)
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		ctx.SetBodyString(`{"error": "Failed to update user"}`)
		return
	}

	log.Printf("DEBUG: User updated successfully with ID: %v", user.ID)
	ctx.SetStatusCode(fasthttp.StatusOK)
	json.NewEncoder(ctx).Encode(user)
}

func (c *UserControllerImpl) DeleteUser(ctx *fasthttp.RequestCtx) {
	log.Printf("DEBUG: DeleteUser endpoint called")

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

type UserRouterImpl struct {
	controller UserController
}

func NewUserRouter(controller UserController) *UserRouterImpl {
	return &UserRouterImpl{controller: controller}
}

func (r *UserRouterImpl) RegisterRoutes(router *fasthttp.RequestHandler) {
	log.Printf("DEBUG: Registering user routes")

	handler := func(ctx *fasthttp.RequestCtx) {
		switch string(ctx.Path()) {
		case "/api/users":
			if ctx.IsPost() {
				r.controller.CreateUser(ctx)
			}
		case "/api/users/:id":
			if ctx.IsGet() {
				r.controller.GetUserByID(ctx)
			} else if ctx.IsPut() {
				r.controller.UpdateUser(ctx)
			} else if ctx.IsDelete() {
				r.controller.DeleteUser(ctx)
			}
		case "/api/users/username/:username":
			if ctx.IsGet() {
				r.controller.GetUserByUsername(ctx)
			}
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(`{"error": "Not found"}`)
		}
	}

	*router = handler
}
