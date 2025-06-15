package basic

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/jwt"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

// IsAuthenticated takes the endpoint context and checks if the user is authenticated.
// It returns a boolean value.
func IsAuthenticated(ctx *fasthttp.RequestCtx) bool {
	var pth = string(ctx.Path())

	authHeaderToken := string(ctx.Request.Header.Peek("Authorization"))

	err, isValid := jwt.VerifyToken(authHeaderToken)
	if err != nil {
		log.Printf("(Error): error verifying token at endpoint (%s).", pth)

		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "5", Values: []string{"Error verifying token: " + err.Error()}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return false
	}

	if !isValid {
		log.Printf("(Error): invalid token at endpoint (%s).", pth)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "5", Values: []string{"Invalid token: " + err.Error()}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return false
	}

	// Get username from token
	username, err := jwt.GetUsernameFromToken(authHeaderToken)
	if err != nil {
		log.Printf("(Error): error getting username from token at endpoint (%s).", pth)
		return false
	}

	if username == "" {
		log.Printf("(Error): error getting username from token at endpoint (%s).", pth)
		return false
	}

	ctx.SetUserValue("username", username)
	return true
}
