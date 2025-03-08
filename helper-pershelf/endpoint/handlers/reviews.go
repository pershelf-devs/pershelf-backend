package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetAllReviewsHandler retrieves all reviews from the database and sends them in a response to the client's request.
func GetAllReviewsHandler(ctx *fasthttp.RequestCtx) {
	var reviews []crud.Review
	if reviews = crud.GetAllReviews(); reviews == nil {
		log.Printf("(Error): error retrieving reviews list at endpoint (%s).", string(ctx.Path()))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): reviews list retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(response.ReviewsResp{
		Status:  response.ResponseMessage{Code: "0", Values: nil},
		Reviews: reviews,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(ctx.Path()))
	}
}

// GetReviewByIDHandler retrieves a review by ID from the database and sends it in the response.
func GetReviewByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth           = ctx.Path()
		reviewID, err = strconv.Atoi(ctx.UserValue("id").(string))
		review        crud.Review
	)

	if err != nil {
		log.Printf("(Error): error converting review ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	review = crud.GetReviewByID(reviewID)
	if review.ID == 0 {
		log.Printf("(Error): review not found by ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): review retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(review); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetReviewsByUserIDHandler retrieves all reviews by user ID.
func GetReviewsByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
		reviews     []crud.Review
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	reviews = crud.GetReviewsByUserID(userID)
	if reviews == nil {
		log.Printf("(Error): no reviews found for user ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): reviews retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(reviews); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetReviewsByBookIDHandler retrieves all reviews for a specific book ID.
func GetReviewsByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("book-id").(string))
		reviews     []crud.Review
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	reviews = crud.GetReviewsByBookID(bookID)
	if reviews == nil {
		log.Printf("(Error): no reviews found for book ID at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusNotFound)
		return
	}

	log.Printf("(Information): reviews retrieved successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
	if err := json.NewEncoder(ctx).Encode(reviews); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateReviewHandler creates a new review in the database.
func CreateReviewHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth    = ctx.Path()
		review crud.Review
	)

	if err := json.Unmarshal(ctx.Request.Body(), &review); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	review = crud.CreateReview(&review)
	if review.ID == 0 {
		log.Printf("(Error): error creating review at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): review created successfully.")
	ctx.SetStatusCode(fasthttp.StatusCreated)
}

// UpdateReviewHandler updates an existing review in the database.
func UpdateReviewHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth    = ctx.Path()
		review crud.Review
	)

	if err := json.Unmarshal(ctx.Request.Body(), &review); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s): %v", pth, err)
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	review = crud.UpdateReview(review)
	if review.ID == 0 {
		log.Printf("(Error): error updating review at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): review updated successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}

// DeleteReviewHandler deletes a review from the database.
func DeleteReviewHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth           = ctx.Path()
		reviewID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting review ID to int at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusBadRequest)
		return
	}

	if err := crud.DeleteReview(reviewID); err != nil {
		log.Printf("(Error): error deleting review at endpoint (%s).", string(pth))
		ctx.SetStatusCode(fasthttp.StatusInternalServerError)
		return
	}

	log.Printf("(Information): review deleted successfully.")
	ctx.SetStatusCode(fasthttp.StatusOK)
}
