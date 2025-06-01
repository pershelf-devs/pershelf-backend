package userLikeUtils

import (
	"encoding/json"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetBookLikesByBookID(bookID int) []tablesModels.BookLike {
	if bookID == 0 {
		log.Printf("Invalid book ID")
		return nil
	}

	jsonData, err := helperContact.HelperRequest("/likes/get/bookID/"+strconv.Itoa(bookID), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return nil
	}

	var likeResp response.BookLikesResp
	if err := json.Unmarshal(jsonData, &likeResp); err != nil {
		log.Printf("Error unmarshalling likes: %v", err)
		return nil
	}

	if likeResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", likeResp.Status.Code)
		return nil
	}

	if len(likeResp.BookLikes) == 0 {
		log.Printf("No likes found for book ID %d", bookID)
		return nil
	}

	return likeResp.BookLikes
}

func GetBookLikesByUserID(userID int) []tablesModels.BookLike {
	if userID == 0 {
		log.Printf("Invalid user ID")
		return nil
	}

	jsonData, err := helperContact.HelperRequest("/likes/get/userID/"+strconv.Itoa(userID), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return nil
	}

	var likeResp response.BookLikesResp
	if err := json.Unmarshal(jsonData, &likeResp); err != nil {
		log.Printf("Error unmarshalling likes: %v", err)
		return nil
	}

	if likeResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", likeResp.Status.Code)
		return nil
	}

	if len(likeResp.BookLikes) == 0 {
		log.Printf("No likes found for user ID %d", userID)
		return nil
	}

	return likeResp.BookLikes
}
