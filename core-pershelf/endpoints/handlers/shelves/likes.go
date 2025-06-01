package shelves

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/shelfBookUtils"
	"github.com/core-pershelf/internal/helperUtils/shelfUtils"
	"github.com/core-pershelf/internal/helperUtils/userShelfUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

type LikeRequest struct {
	UserID     int  `json:"user_id"`
	BookID     int  `json:"book_id"`
	LikeStatus bool `json:"like_status"` // true: like, false: unlike
}

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
		var req LikeRequest
		if err := json.Unmarshal(ctx.Request.Body(), &req); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			log.Printf("Request body: %s", string(ctx.Request.Body()))
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		// 1. Ensure "likes" shelf exists, if not create it
		userShelf := shelfUtils.GetShelfByUserIDAndName(req.UserID, "likes")
		if userShelf.ID == 0 {
			tempShelf := tablesModels.UserShelf{
				UserID:    req.UserID,
				ShelfName: "likes",
			}
			newShelf, err := userShelfUtils.CreateUserShelf(tempShelf)
			if err != nil {
				log.Printf("Error creating likes shelf: %v", err)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error creating likes shelf"}}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}
			userShelf = newShelf
		}

		// 2. Remove book from "likes" shelf if like status is false
		if !req.LikeStatus {
			if err := shelfBookUtils.DeleteShelfBookByShelfIDAndBookID(userShelf.ID, req.BookID); err != nil {
				log.Printf("Error deleting shelf book: %v", err)
				if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error deleting shelf book"}}); err != nil {
					log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
				}
				return
			}
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: []string{"Book unliked"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		// 3. Add book to "likes" shelf if like status is true
		tempShelfBook := tablesModels.ShelfBook{
			ShelfID: userShelf.ID,
			BookID:  req.BookID,
		}

		newShelfBook, err := shelfBookUtils.CreateShelfBook(tempShelfBook)
		if err != nil {
			log.Printf("Error creating shelf book: %v", err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error creating shelf book"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if newShelfBook.ID == 0 {
			log.Printf("Error creating shelf book: %v", err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error creating shelf book"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "0", Values: []string{"Book liked"}}); err != nil {
			log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
		}
	}
}
