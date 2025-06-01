package reviews

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/reviewUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

func GetReviewsByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var bookID int
		if err := json.Unmarshal(ctx.Request.Body(), &bookID); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			log.Printf("Request body: %s", string(ctx.Request.Body()))
			if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		reviews, err := reviewUtils.GetReviewsByBookID(bookID)
		if err != nil {
			log.Printf("Error retrieving reviews by book ID at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error retrieving reviews by book ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		log.Printf("Reviews retrieved successfully at endpoint (%s).", pth)
		if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
			Status:  response.ResponseMessage{Code: "0", Values: nil},
			Reviews: reviews,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return
	}
}

func GetReviewsByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var UserID int
		if err := json.Unmarshal(ctx.Request.Body(), &UserID); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			log.Printf("Request body: %s", string(ctx.Request.Body()))
			if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		reviews, err := reviewUtils.GetReviewsByUserID(UserID)
		if err != nil {
			log.Printf("Error getting reviews by user ID at endpoint (%s): %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error retrieving reviews by book ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
			Status:  response.ResponseMessage{Code: "0", Values: nil},
			Reviews: reviews,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		} else {
			log.Printf("Reviews retrieved successfully at endpoint (%s).", pth)
		}

	}
}
