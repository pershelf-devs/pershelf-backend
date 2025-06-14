package reviewUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	customizedmodels "github.com/core-pershelf/rest/helperContact/customizedModels"
	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

// GetDetailedReviewsByBookID gets detailed reviews by book ID
func GetDetailedReviewsByBookID(bookID int) ([]customizedmodels.DetailedReview, error) {
	jsonData, err := helperContact.HelperRequest("/reviews/get/detailed-reviews/book-id/"+strconv.Itoa(bookID), nil)
	if err != nil {
		log.Printf("Error getting detailed reviews by book ID %d: %v", bookID, err)
		return nil, fmt.Errorf("error getting detailed reviews by book ID: %v", err)
	}

	var detailedReviewsResp response.DetailedReviewsResp
	if err := json.Unmarshal(jsonData, &detailedReviewsResp); err != nil {
		log.Printf("Error unmarshalling detailed reviews: %v", err)
		return nil, fmt.Errorf("error unmarshalling detailed reviews: %v", err)
	}

	if detailedReviewsResp.Status.Code != "0" {
		return nil, fmt.Errorf("error getting detailed reviews by book ID: %s", detailedReviewsResp.Status.Code)
	}

	return detailedReviewsResp.DetailedReviews, nil
}
