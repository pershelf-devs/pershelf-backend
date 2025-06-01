package reviews

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/reviewUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

// CreateBookReviewHandler creates a new book review
func CreateBookReviewHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var review tablesModels.Review
		if err := json.Unmarshal(ctx.Request.Body(), &review); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			log.Printf("Request body: %s", string(ctx.Request.Body()))
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "3",
				Values: []string{"Error unmarshalling the request body"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Check if the book ID is provided
		if review.BookID == 0 {
			log.Printf("Book ID is required at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Book ID is required"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
		}

		// Check if the user ID is provided
		if review.UserID == 0 {
			log.Printf("User ID is required at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"User ID is required"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Create the review
		if _, err := reviewUtils.CreateReview(review); err != nil {
			log.Printf("Error creating review at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "3",
				Values: []string{"Error creating review"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Success
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "0",
			Values: []string{"Review created successfully"},
		}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
		return
	}
}
