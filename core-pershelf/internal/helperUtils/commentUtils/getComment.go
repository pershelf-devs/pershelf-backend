package commentUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllComments() ([]tablesModels.Comment, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/comments/get/all", nil)
	if err != nil {
		log.Printf("Error getting all comments: %v", err)
		return nil, fmt.Errorf("error getting all comments: %v", err)
	}

	var commentsResp response.CommentsResp
	if err := json.Unmarshal(jsonData, &commentsResp); err != nil {
		log.Printf("Error unmarshalling comments: %v", err)
		return nil, fmt.Errorf("error unmarshalling comments: %v", err)
	}

	if commentsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", commentsResp.Status.Code)
		return nil, fmt.Errorf("error getting comments: %s", commentsResp.Status.Code)
	}

	return commentsResp.Comments, nil
}

func GetCommentByID(id int) (tablesModels.Comment, error) {
	if id == 0 {
		return tablesModels.Comment{}, fmt.Errorf("invalid comment ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/comments/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting comment by ID %d: %v", id, err)
		return tablesModels.Comment{}, fmt.Errorf("error getting comment: %v", err)
	}

	var commentsResp response.CommentsResp
	if err := json.Unmarshal(jsonData, &commentsResp); err != nil {
		log.Printf("Error unmarshalling comment: %v", err)
		return tablesModels.Comment{}, fmt.Errorf("error unmarshalling comment: %v", err)
	}

	if commentsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", commentsResp.Status.Code)
		return tablesModels.Comment{}, fmt.Errorf("error getting comment: %s", commentsResp.Status.Code)
	}

	if len(commentsResp.Comments) == 0 {
		log.Printf("No comment found with ID %d", id)
		return tablesModels.Comment{}, fmt.Errorf("comment not found")
	}

	return commentsResp.Comments[0], nil
}
