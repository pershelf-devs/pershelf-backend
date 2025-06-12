package userBookRelations

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/userBookRelationUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

// LikeBookHandler handles the user's request to like or dislike a book
func LikeBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var likeBookRequest LikeBookRequest
		if err := json.Unmarshal(ctx.Request.Body(), &likeBookRequest); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			log.Printf("Request body: %s", string(ctx.Request.Body()))
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if likeBookRequest.UserID <= 0 {
			log.Printf("Invalid user ID (retrieved: %d) at endpoint (%s).", likeBookRequest.UserID, pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if likeBookRequest.BookID <= 0 {
			log.Printf("Invalid book ID (retrieved: %d) at endpoint (%s).", likeBookRequest.BookID, pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		// Get user book relation by user ID and book ID
		userBookRelation := userBookRelationUtils.GetUserBookRelationByUserIDAndBookID(likeBookRequest.UserID, likeBookRequest.BookID)
		userBookRelation.Like = !userBookRelation.Like
		userBookRelation.UserID = likeBookRequest.UserID
		userBookRelation.BookID = likeBookRequest.BookID

		if userBookRelation.ID == 0 {
			// Create user book relation
			userBookRelation = userBookRelationUtils.CreateUserBookRelation(userBookRelation)
			if userBookRelation.ID == 0 {
				log.Printf("Error creating user book relation at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error creating user book relation"}}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}

			statusCode := "100"
			if !userBookRelation.Like {
				statusCode = "101"
			}

			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: statusCode, Values: []string{"Successfully updated book like status"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		// Update user book relation
		code, err := userBookRelationUtils.UpdateUserBookRelation(userBookRelation)
		if err != nil {
			log.Printf("Error updating user book relation at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if code != "0" {
			log.Printf("Error updating user book relation at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		statusCode := "100"
		if !userBookRelation.Like {
			statusCode = "101"
		}

		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: statusCode, Values: []string{"Successfully updated book like status"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return
	}
}
