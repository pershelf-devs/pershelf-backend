package auth

import (
	"context"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/core-pershelf/globals"
	"github.com/core-pershelf/jwt"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func ClassicAuthHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth           = string(ctx.Path())
		receivedCreds struct {
			Email    string `json:"email"`
			Password string `json:"password"`
		}
	)

	// decode received credentials (username, password, client identifier)
	if err := json.Unmarshal(ctx.Request.Body(), &receivedCreds); err != nil {
		log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// check if the received credentials are valid
	if receivedCreds.Email == "" || receivedCreds.Password == "" {
		log.Printf("Invalid credentials at endpoint %s", pth)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Find the user from Mongo db
	var user tablesModels.User
	filter := bson.M{"email": receivedCreds.Email}
	err := globals.UsersCollection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			log.Printf("User with email (%s) not found", receivedCreds.Email)
			if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "10",
				Values: []string{receivedCreds.Email},
			}); err != nil {
				log.Printf("Error encoding response message: %v", err)
			}
			return
		}
		log.Printf("Error finding user: %v", err)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "3",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Hash the password
	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(receivedCreds.Password)))

	// Debug
	log.Printf("HashedPass %v", hashedPassword)
	log.Printf("Correct hash %v", user.Password)

	// Check if the password is correct
	if user.Password != hashedPassword {
		log.Printf("Invalid password for user with email (%s)", receivedCreds.Email)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "11",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Generate a token
	accessToken, err := jwt.CreateJwtToken(user.Username, 15*time.Minute)
	if err != nil {
		log.Printf("Error creating JWT token: %v", err)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "12",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Generate a refresh token
	refreshToken, err := jwt.CreateJwtToken(user.Username, 7*24*time.Hour)
	if err != nil {
		log.Printf("Error creating JWT token: %v", err)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "12",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	type TempLoginStruct struct {
		UserInfo tablesModels.User `json:"userInfo"`
		Token    string            `json:"token"`
		Refresh  string            `json:"refresh"`
	}

	type TempLoginResp struct {
		Status response.ResponseMessage `json:"status"`
		Data   TempLoginStruct          `json:"data"`
	}

	// remove the password from the user info
	user.Password = ""

	// Return the token
	if err = json.NewEncoder(ctx).Encode(TempLoginResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Data: TempLoginStruct{
			UserInfo: user,
			Token:    accessToken,
			Refresh:  refreshToken,
		},
	}); err != nil {
		log.Printf("Error encoding response message: %v", err)
	}
}
