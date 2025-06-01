package userLikeUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteBookLikeByID(id int) error {
	if id == 0 {
		log.Printf("Invalid book like ID")
		return fmt.Errorf("invalid book like ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/likes/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting book like: %v", err)
	}

	var likeResp response.BookLikesResp
	if err := json.Unmarshal(jsonData, &likeResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if likeResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", likeResp.Status.Code)
		return fmt.Errorf("error deleting book like: %s", likeResp.Status.Code)
	}

	return nil
}

func DeleteBookLikesByBookIDAndUserID(userID int, bookID int) error {
	if userID == 0 || bookID == 0 {
		log.Printf("Invalid user ID or book ID")
		return fmt.Errorf("invalid user ID or book ID")
	}

	url := "/likes/delete/userID/" + strconv.Itoa(userID) + "/bookID/" + strconv.Itoa(bookID)
	jsonData, err := helperContact.HelperRequest(url, nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting book like: %v", err)
	}

	var likeResp response.BookLikesResp
	if err := json.Unmarshal(jsonData, &likeResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if likeResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", likeResp.Status.Code)
		return fmt.Errorf("error deleting book like: %s", likeResp.Status.Code)
	}

	return nil
}
