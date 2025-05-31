package books

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

func GetUserRecomendedBooksHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: nil}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
		return

	default:
		var request struct {
			UserID string `json:"user_id"`
		}

		if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling the request body. No user id provided"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// check if the user id is provided
		if request.UserID == "" {
			log.Printf("User ID is required at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"User ID is required"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// TODO: Implement the logic to get the user recomended books
	}
}
