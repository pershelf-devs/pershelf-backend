package commentUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func CreateComment(comment tablesModels.Comment) (tablesModels.Comment, error) {
	// Marshal the comment into JSON
	payload, err := json.Marshal(comment)
	if err != nil {
		log.Printf("Error marshalling comment: %v", err)
		return tablesModels.Comment{}, fmt.Errorf("error marshalling comment: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/comments/create", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.Comment{}, fmt.Errorf("error creating comment: %v", err)
	}

	var commentsResp response.CommentsResp
	if err := json.Unmarshal(jsonData, &commentsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.Comment{}, fmt.Errorf("error unmarshalling created comment: %v", err)
	}

	if commentsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", commentsResp.Status.Code)
		return tablesModels.Comment{}, fmt.Errorf("error creating comment: %s", commentsResp.Status.Code)
	}

	if len(commentsResp.Comments) == 0 {
		log.Printf("Comment creation succeeded but no comment returned")
		return tablesModels.Comment{}, fmt.Errorf("no comment returned after creation")
	}

	return commentsResp.Comments[0], nil
}
