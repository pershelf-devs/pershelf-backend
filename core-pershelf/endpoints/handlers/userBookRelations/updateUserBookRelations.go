package userBookRelations

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/userBookRelationUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
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
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if likeBookRequest.UserID <= 0 {
			log.Printf("Invalid user ID (retrieved: %d) at endpoint (%s).", likeBookRequest.UserID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if likeBookRequest.BookID <= 0 {
			log.Printf("Invalid book ID (retrieved: %d) at endpoint (%s).", likeBookRequest.BookID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}},
			}); err != nil {
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
				if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
					Status: response.ResponseMessage{Code: "3", Values: []string{"Error creating user book relation"}},
				}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}
		} else {
			// Update user book relation
			code, err := userBookRelationUtils.UpdateUserBookRelation(userBookRelation)
			if err != nil {
				log.Printf("Error updating user book relation at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
					Status: response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}},
				}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}

			if code != "0" {
				log.Printf("Error updating user book relation at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
					Status: response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}},
				}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}
		}

		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "0", Values: []string{"Successfully updated book like status"}},
			UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return
	}
}

// FavoriteBookHandler handles the user's request to favorite or unfavorite a book
func FavoriteBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var requestBody LikeBookRequest
		if err := json.Unmarshal(ctx.Request.Body(), &requestBody); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if requestBody.UserID <= 0 {
			log.Printf("Invalid user ID (retrieved: %d) at endpoint (%s).", requestBody.UserID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if requestBody.BookID <= 0 {
			log.Printf("Invalid book ID (retrieved: %d) at endpoint (%s).", requestBody.BookID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		userBookRelation := userBookRelationUtils.GetUserBookRelationByUserIDAndBookID(requestBody.UserID, requestBody.BookID)
		userBookRelation.Favorite = !userBookRelation.Favorite
		userBookRelation.UserID = requestBody.UserID
		userBookRelation.BookID = requestBody.BookID

		if userBookRelation.ID == 0 {
			userBookRelation = userBookRelationUtils.CreateUserBookRelation(userBookRelation)
			if userBookRelation.ID == 0 {
				log.Printf("Error creating user book relation at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error creating user book relation"}}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}

			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status:            response.ResponseMessage{Code: "0", Values: []string{"Successfully created user book relation"}},
				UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

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

		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "0", Values: []string{"Successfully updated user book relation"}},
			UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
	}
}

// AddToReadListHandler handles the user's request to add a book to the read list
func AddToReadListHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var requestBody LikeBookRequest
		if err := json.Unmarshal(ctx.Request.Body(), &requestBody); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if requestBody.UserID <= 0 {
			log.Printf("Invalid user ID (retrieved: %d) at endpoint (%s).", requestBody.UserID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if requestBody.BookID <= 0 {
			log.Printf("Invalid book ID (retrieved: %d) at endpoint (%s).", requestBody.BookID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		userBookRelation := userBookRelationUtils.GetUserBookRelationByUserIDAndBookID(requestBody.UserID, requestBody.BookID)
		userBookRelation.ReadList = !userBookRelation.ReadList
		userBookRelation.UserID = requestBody.UserID
		userBookRelation.BookID = requestBody.BookID

		// Create user book relation if it doesn't exist
		if userBookRelation.ID == 0 {
			userBookRelation = userBookRelationUtils.CreateUserBookRelation(userBookRelation)
			if userBookRelation.ID == 0 {
				log.Printf("Error creating user book relation at endpoint (%s).", pth)
				if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
					Status: response.ResponseMessage{Code: "3", Values: []string{"Error creating user book relation"}},
				}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}

			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status:            response.ResponseMessage{Code: "0", Values: []string{"Successfully created user book relation"}},
				UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		code, err := userBookRelationUtils.UpdateUserBookRelation(userBookRelation)
		if err != nil {
			log.Printf("Error updating user book relation at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if code != "0" {
			log.Printf("Error updating user book relation at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: code, Values: []string{"Error updating user book relation"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "0", Values: []string{"Successfully updated user book relation"}},
			UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return
	}
}

// SetAsReadHandler handles the user's request to set a book as read
func SetAsReadHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var requestBody LikeBookRequest
		if err := json.Unmarshal(ctx.Request.Body(), &requestBody); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if requestBody.UserID <= 0 {
			log.Printf("Invalid user ID (retrieved: %d) at endpoint (%s).", requestBody.UserID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid user ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if requestBody.BookID <= 0 {
			log.Printf("Invalid book ID (retrieved: %d) at endpoint (%s).", requestBody.BookID, pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		userBookRelation := userBookRelationUtils.GetUserBookRelationByUserIDAndBookID(requestBody.UserID, requestBody.BookID)
		userBookRelation.Read = !userBookRelation.Read
		userBookRelation.UserID = requestBody.UserID
		userBookRelation.BookID = requestBody.BookID

		// If the book is not read,
		if !userBookRelation.Read {
			userBookRelation.Favorite = false
			userBookRelation.Like = false
		} else {
			// If the book is read
			userBookRelation.ReadList = false
		}

		code, err := userBookRelationUtils.UpdateUserBookRelation(userBookRelation)
		if err != nil {
			log.Printf("Error updating user book relation at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error updating user book relation"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if code != "0" {
			log.Printf("Error updating user book relation at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
				Status: response.ResponseMessage{Code: code, Values: []string{"Error updating user book relation"}},
			}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "0", Values: []string{"Successfully updated user book relation"}},
			UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return
	}
}
