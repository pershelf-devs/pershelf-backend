package users

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/userUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

func GetUserByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = string(ctx.Path())
		user tablesModels.User
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var userID int
		if err := json.Unmarshal(ctx.Request.Body(), &userID); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Get the user by ID
		if user = userUtils.GetUserByID(userID); user.ID == 0 {
			log.Printf("Error getting user by ID at endpoint %s: %v", pth, user)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error getting user by ID"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Prepare the response
		user.Password = "" // Ensure password is not sent in the response
		if err := json.NewEncoder(ctx).Encode(response.UsersResp{
			Status: response.ResponseMessage{Code: "0"},
			Users:  []tablesModels.User{user},
		}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			ctx.SetStatusCode(fasthttp.StatusInternalServerError)
			return
		}
	}
}
