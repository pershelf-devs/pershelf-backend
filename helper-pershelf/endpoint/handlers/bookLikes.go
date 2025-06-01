package handlers

import (
	"encoding/json"
	"log"
	"strconv"

	"github.com/pershelf/pershelf/crud"
	"github.com/pershelf/pershelf/endpoint/response"
	"github.com/valyala/fasthttp"
)

// GetBookLikesByBookIDHandler retrieves all book likes by book ID
func GetBookLikesByBookIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		bookID, err = strconv.Atoi(ctx.UserValue("book-id").(string))
		bookLikes   []crud.BookLike
	)

	if err != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookID == 0 {
		log.Printf("(Error): book ID is 0 at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Book ID is 0"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookLikes = crud.GetBookLikesByBookID(bookID); bookLikes == nil {
		log.Printf("(Error): error retrieving book likes by book ID at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Error retrieving book likes by book ID"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book likes retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		BookLikes: bookLikes,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// GetBookLikesByUserIDHandler retrieves all book likes by user ID
func GetBookLikesByUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth         = ctx.Path()
		userID, err = strconv.Atoi(ctx.UserValue("user-id").(string))
		bookLikes   []crud.BookLike
	)

	if err != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userID == 0 {
		log.Printf("(Error): user ID is 0 at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"User ID is 0"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookLikes = crud.GetBookLikesByUserID(userID); bookLikes == nil {
		log.Printf("(Error): error retrieving book likes by user ID at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Error retrieving book likes by user ID"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book likes retrieved successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		BookLikes: bookLikes,
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// CreateBookLikeHandler creates a new book like
func CreateBookLikeHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth      = ctx.Path()
		bookLike crud.BookLike
	)

	if err := json.Unmarshal(ctx.Request.Body(), &bookLike); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.CreateBookLike(&bookLike); err != nil {
		log.Printf("(Error): error creating book like at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
			Status:    response.ResponseMessage{Code: "3", Values: []string{"Error creating book like"}},
			BookLikes: nil,
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book like created successfully.")
	if err := json.NewEncoder(ctx).Encode(response.BookLikesResp{
		Status:    response.ResponseMessage{Code: "0", Values: nil},
		BookLikes: []crud.BookLike{bookLike},
	}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// UpdateBookLikeHandler updates a book like
func UpdateBookLikeHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth      = ctx.Path()
		bookLike crud.BookLike
	)

	if err := json.Unmarshal(ctx.Request.Body(), &bookLike); err != nil {
		log.Printf("(Error): error unmarshalling request body at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.UpdateBookLike(bookLike); err != nil {
		log.Printf("(Error): error updating book like at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error updating book like"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
	}

	log.Printf("(Information): book like updated successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteBookLikeByIDHandler deletes a book like
func DeleteBookLikeByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth             = ctx.Path()
		bookLikeID, err = strconv.Atoi(ctx.UserValue("id").(string))
	)

	if err != nil {
		log.Printf("(Error): error converting book like ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting book like ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookLikeID == 0 {
		log.Printf("(Error): book like ID is 0 at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Book like ID is 0"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteBookLikeByID(bookLikeID); err != nil {
		log.Printf("(Error): error deleting book like at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting book like"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book like deleted successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}

// DeleteBookLikesByBookIDAndUserIDHandler deletes a book like by book ID and user ID
func DeleteBookLikesByBookIDAndUserIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth          = ctx.Path()
		bookID, err1 = strconv.Atoi(ctx.UserValue("book-id").(string))
		userID, err2 = strconv.Atoi(ctx.UserValue("user-id").(string))
	)

	if err1 != nil {
		log.Printf("(Error): error converting book ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting book ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err2 != nil {
		log.Printf("(Error): error converting user ID to int at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error converting user ID to int"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if bookID == 0 {
		log.Printf("(Error): book ID is 0 at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Book ID is 0"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if userID == 0 {
		log.Printf("(Error): user ID is 0 at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"User ID is 0"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	if err := crud.DeleteBookLikesByBookIDAndUserID(bookID, userID); err != nil {
		log.Printf("(Error): error deleting book likes by book ID and user ID at endpoint (%s).", string(pth))
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting book likes by book ID and user ID"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
		}
		return
	}

	log.Printf("(Information): book likes deleted successfully.")
	if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: nil}); err != nil {
		log.Printf("(Error): error encoding response message at endpoint (%s).", string(pth))
	}
}
