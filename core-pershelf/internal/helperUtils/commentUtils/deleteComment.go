package commentUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteComment(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid comment ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/comments/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting comment: %v", err)
	}

	var commentsResp response.CommentsResp
	if err := json.Unmarshal(jsonData, &commentsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if commentsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", commentsResp.Status.Code)
		return fmt.Errorf("error deleting comment: %s", commentsResp.Status.Code)
	}

	return nil
}
