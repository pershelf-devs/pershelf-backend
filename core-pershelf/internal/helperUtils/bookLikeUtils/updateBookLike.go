package userLikeUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateBookLike(bookLike tablesModels.BookLike) (tablesModels.BookLike, error) {
	// Marshal the bookLike into JSON
	payload, err := json.Marshal(bookLike)
	if err != nil {
		log.Printf("Error marshalling book like: %v", err)
		return tablesModels.BookLike{}, fmt.Errorf("error marshalling book like: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/likes/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.BookLike{}, fmt.Errorf("error updating book like: %v", err)
	}

	var likeResp response.BookLikesResp
	if err := json.Unmarshal(jsonData, &likeResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.BookLike{}, fmt.Errorf("error unmarshalling updated book like: %v", err)
	}

	if likeResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", likeResp.Status.Code)
		return tablesModels.BookLike{}, fmt.Errorf("error updating book like: %s", likeResp.Status.Code)
	}

	if len(likeResp.BookLikes) == 0 {
		log.Printf("Book like update succeeded but no book like returned")
		return tablesModels.BookLike{}, fmt.Errorf("no book like returned after update")
	}

	return likeResp.BookLikes[0], nil
}
