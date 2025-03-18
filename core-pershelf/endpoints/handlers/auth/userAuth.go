package auth

import (
	"crypto/md5"
	"encoding/json"
	"log"
	"time"

	"github.com/core-pershelf/jwt"
	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
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

	// Get user by email
	jsonData, err := helperContact.HelperRequest("/users/get/email/"+receivedCreds.Email, nil)
	if err != nil {
		log.Printf("Error getting user by email: %v", err)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "3",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Unmarshal the response
	var usersResp response.UsersResp
	if err := json.Unmarshal(jsonData, &usersResp); err != nil {
		log.Printf("Error unmarshalling user: %v", err)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "3",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// If user not found
	if len(usersResp.Users) == 0 {
		log.Printf("User with email (%s) not found", receivedCreds.Email)
		if err = json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "10",
			Values: []string{receivedCreds.Email},
		}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Hash the password
	hashedPassword := string(md5.New().Sum([]byte(receivedCreds.Password)))

	// Check if the password is correct
	if usersResp.Users[0].Password != hashedPassword {
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
	accessToken, err := jwt.CreateJwtToken(usersResp.Users[0].Username, 15*time.Minute)
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
	refreshToken, err := jwt.CreateJwtToken(usersResp.Users[0].Username, 7*24*time.Hour)
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
	usersResp.Users[0].Password = ""

	// Return the token
	if err = json.NewEncoder(ctx).Encode(TempLoginResp{
		Status: response.ResponseMessage{Code: "0", Values: nil},
		Data: TempLoginStruct{
			UserInfo: usersResp.Users[0],
			Token:    accessToken,
			Refresh:  refreshToken,
		},
	}); err != nil {
		log.Printf("Error encoding response message: %v", err)
	}
}
