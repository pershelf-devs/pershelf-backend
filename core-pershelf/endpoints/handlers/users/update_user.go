package users

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/userUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

func UpdateUserProfilePhotoHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var request struct {
			UserID      int    `json:"user_id"`
			ImageBase64 string `json:"image_base64"`
		}

		if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
			log.Printf("Error unmarshling request body at endpoint (%s)", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		if request.ImageBase64 == "" {
			log.Printf(" invalid request. imagebase64 not found(%s) at endpoint (%s)", request.ImageBase64, pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		user := userUtils.GetUserByID(request.UserID)
		if user.ID == 0 {
			log.Printf("user with ID (%d) not found at endpoint (%s)", request.UserID, pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		user.ImageBase64 = request.ImageBase64

		if err := userUtils.UpdateUser(user); err != nil {
			log.Printf("error updating user at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "4"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: []string{"user imaage updated succesfully"}}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
	}
}
