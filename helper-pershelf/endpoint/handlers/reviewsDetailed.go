package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud/customized"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetDetailedReviewsByBookIDHandler gets detailed reviews by book ID
func GetDetailedReviewsByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth             = ctx.Path()
		bookID, err     = strconv.Atoi(ctx.UserValue("book-id").(string))
		reviewsDetailed = []customized.DetailedReview{}
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.DetailedReviewsResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if reviewsDetailed, err = customized.GetDetailedReviewsByBookID(bookID); err != nil {
		log.Printf("(Error): error getting detailed reviews by book ID at endpoint (%s). %v", string(pth), err)
		if err := json.NewEncoder(ctx).Encode(response.DetailedReviewsResp{
			Status: response.ResponseMessage{Code: "3", Values: []string{"Error getting detailed reviews by book ID"}},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): detailed reviews retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.DetailedReviewsResp{
		Status:          response.ResponseMessage{Code: "0", Values: nil},
		DetailedReviews: reviewsDetailed,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}
