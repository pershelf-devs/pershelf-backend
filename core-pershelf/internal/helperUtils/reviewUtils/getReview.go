package reviewUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllReviews() ([]tablesModels.Review, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/reviews/get/all", nil)
	if err != nil {
		log.Printf("Error getting all reviews: %v", err)
		return nil, fmt.Errorf("error getting all reviews: %v", err)
	}

	var reviewsResp response.ReviewsResp
	if err := json.Unmarshal(jsonData, &reviewsResp); err != nil {
		log.Printf("Error unmarshalling reviews: %v", err)
		return nil, fmt.Errorf("error unmarshalling reviews: %v", err)
	}

	if reviewsResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", reviewsResp.Status.Code)
		return nil, fmt.Errorf("error getting reviews: %s", reviewsResp.Status.Code)
	}

	return reviewsResp.Reviews, nil
}

func GetReviewByID(id int) (tablesModels.Review, error) {
	if id == 0 {
		return tablesModels.Review{}, fmt.Errorf("invalid review ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/reviews/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting review by ID %d: %v", id, err)
		return tablesModels.Review{}, fmt.Errorf("error getting review: %v", err)
	}

	var reviewsResp response.ReviewsResp
	if err := json.Unmarshal(jsonData, &reviewsResp); err != nil {
		log.Printf("Error unmarshalling review: %v", err)
		return tablesModels.Review{}, fmt.Errorf("error unmarshalling review: %v", err)
	}

	if reviewsResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", reviewsResp.Status.Code)
		return tablesModels.Review{}, fmt.Errorf("error getting review: %s", reviewsResp.Status.Code)
	}

	if len(reviewsResp.Reviews) == 0 {
		log.Printf("No review found with ID %d", id)
		return tablesModels.Review{}, fmt.Errorf("review not found")
	}

	return reviewsResp.Reviews[0], nil
}

func GetReviewsByUserID(userID int) ([]tablesModels.Review, error) {
	if userID == 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/reviews/get/user-id/"+strconv.Itoa(userID), nil)
	if err != nil {
		log.Printf("Error getting reviews by user ID %d: %v", userID, err)
		return nil, fmt.Errorf("error getting reviews: %v", err)
	}

	var reviewsResp response.ReviewsResp
	if err := json.Unmarshal(jsonData, &reviewsResp); err != nil {
		log.Printf("Error unmarshalling reviews: %v", err)
		return nil, fmt.Errorf("error unmarshalling reviews: %v", err)
	}

	if reviewsResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", reviewsResp.Status.Code)
		return nil, fmt.Errorf("error getting reviews: %s", reviewsResp.Status.Code)
	}

	return reviewsResp.Reviews, nil
}

func GetReviewsByBookID(bookID int) ([]tablesModels.Review, error) {
	if bookID == 0 {
		return nil, fmt.Errorf("invalid book ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/reviews/get/book-id/"+strconv.Itoa(bookID), nil)
	if err != nil {
		log.Printf("Error getting reviews by book ID %d: %v", bookID, err)
		return nil, fmt.Errorf("error getting reviews: %v", err)
	}

	var reviewsResp response.ReviewsResp
	if err := json.Unmarshal(jsonData, &reviewsResp); err != nil {
		log.Printf("Error unmarshalling reviews: %v", err)
		return nil, fmt.Errorf("error unmarshalling reviews: %v", err)
	}

	if reviewsResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", reviewsResp.Status.Code)
		return nil, fmt.Errorf("error getting reviews: %s", reviewsResp.Status.Code)
	}

	return reviewsResp.Reviews, nil
}
