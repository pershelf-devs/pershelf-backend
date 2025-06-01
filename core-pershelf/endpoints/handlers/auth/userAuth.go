package auth

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/core-pershelf/internal/helperUtils/userUtils"
	"github.com/core-pershelf/jwt"
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

	//get user by e-mail
	user := userUtils.GetUserByEmail(receivedCreds.Email)
	if user.ID == 0 {
		log.Printf("e-mail received from user %+v", user)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"User not found"}}); err != nil {
			log.Printf("Error encoding response message: %v", err)
		}
		return
	}

	// Hash the password
	hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(receivedCreds.Password)))

	// Check if the password is correct
	if user.Password != hashedPassword {
		log.Printf("Invalid password for user with email (%s)", receivedCreds.Email)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
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
