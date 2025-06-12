package userBookRelations

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/userBookRelationUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

type LikeBookRequest struct {
	UserID int `json:"user_id"`
	BookID int `json:"book_id"`
}

// GetUserBookRelationByUserIDAndBookIDHandler gets a user book relation by user id and book id
func GetUserBookRelationByUserIDAndBookIDHandler(ctx *fasthttp.RequestCtx) {
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

		userBookRelation := userBookRelationUtils.GetUserBookRelationByUserIDAndBookID(likeBookRequest.UserID, likeBookRequest.BookID)
		if userBookRelation.ID == 0 {
			log.Printf("User book relation not found at endpoint (%s).", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"User book relation not found"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.UserBookRelationsResp{
			Status:            response.ResponseMessage{Code: "0", Values: nil},
			UserBookRelations: []tablesModels.UserBookRelation{userBookRelation},
		}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
		return
	}
}
