package reviewUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteReview(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid review ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/reviews/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting review: %v", err)
	}

	var reviewsResp response.ReviewsResp
	if err := json.Unmarshal(jsonData, &reviewsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if reviewsResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", reviewsResp.Status.Code)
		return fmt.Errorf("error deleting review: %s", reviewsResp.Status.Code)
	}

	return nil
}
