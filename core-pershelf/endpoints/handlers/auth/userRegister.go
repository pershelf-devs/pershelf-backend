package auth

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/core-pershelf/globals"
	"github.com/core-pershelf/mongo/tablesModels/users"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type UserRegisterRequest struct {
	Name            string `json:"name"`
	Surname         string `json:"surname"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	Password        string `json:"password"`
	ConfirmPassword string `json:"confirm_password"`
}

func UserRegisterHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var request UserRegisterRequest

		// Unmarshal the request body
		if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Check if the password and confirm password match
		if request.Password != request.ConfirmPassword {
			log.Printf("Error: Password and ConfirmPassword do not match at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "xxxx", Values: nil}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// hash the password
		hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(request.Password)))

		// Create a new user
		user := users.User{
			ID:        primitive.NewObjectID(),
			Username:  request.Username,
			Password:  hashedPassword,
			Email:     request.Email,
			Name:      request.Name,
			Surname:   request.Surname,
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
		}

		// Initialize user service
		userService := users.NewUserService(users.NewUserRepositoryMongo(globals.UsersCollection))

		// Save the user using the service
		if err := userService.CreateUser(&user); err != nil {
			log.Printf("Error saving the user to the database at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Return the user
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: []string{"User registered successfully"}}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
	}
}
