package reviewUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateReview(review tablesModels.Review) (tablesModels.Review, error) {
	// Marshal the review into JSON
	payload, err := json.Marshal(review)
	if err != nil {
		log.Printf("Error marshalling review: %v", err)
		return tablesModels.Review{}, fmt.Errorf("error marshalling review: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/reviews/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.Review{}, fmt.Errorf("error updating review: %v", err)
	}

	var reviewsResp response.ReviewsResp
	if err := json.Unmarshal(jsonData, &reviewsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.Review{}, fmt.Errorf("error unmarshalling updated review: %v", err)
	}

	if reviewsResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", reviewsResp.Status.Code)
		return tablesModels.Review{}, fmt.Errorf("error updating review: %s", reviewsResp.Status.Code)
	}

	if len(reviewsResp.Reviews) == 0 {
		log.Printf("Review update succeeded but no review returned")
		return tablesModels.Review{}, fmt.Errorf("no review returned after update")
	}

	return reviewsResp.Reviews[0], nil
}
