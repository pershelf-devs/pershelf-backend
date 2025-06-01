package auth

import (
	"crypto/md5"
	"encoding/json"
	"fmt"
	"log"

	"github.com/core-pershelf/internal/helperUtils/userUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
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
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// hash the password
		hashedPassword := fmt.Sprintf("%x", md5.Sum([]byte(request.Password)))

		// user struct
		user := tablesModels.User{
			Username: request.Username,
			Surname:  request.Surname,
			Name:     request.Name,
			Email:    request.Email,
			Password: hashedPassword,
		}
		//create user
		user, err := userUtils.CreateUser(user)
		if err != nil {
			log.Printf("error creating a user at endpoint %s : %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}
		//eğer oluşturulan userın ID si 0 sa hata döndür
		if user.ID == 0 {
			log.Printf("creation fault at endpoint %s : %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
				log.Printf("user is not equal to 0 %s: %v", pth, err)
			}
			return
		}
		// Return the user
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: []string{"User registered successfully"}}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
	}
}
