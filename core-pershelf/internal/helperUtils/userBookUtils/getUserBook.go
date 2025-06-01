package userBookUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllUserBooks() ([]tablesModels.UserBook, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/userBooks/get/all", nil)
	if err != nil {
		log.Printf("Error getting all user books: %v", err)
		return nil, fmt.Errorf("error getting all user books: %v", err)
	}

	var userBooksResp response.UserBooksResp
	if err := json.Unmarshal(jsonData, &userBooksResp); err != nil {
		log.Printf("Error unmarshalling user books: %v", err)
		return nil, fmt.Errorf("error unmarshalling user books: %v", err)
	}

	if userBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", userBooksResp.Status.Code)
		return nil, fmt.Errorf("error getting user books: %s", userBooksResp.Status.Code)
	}

	return userBooksResp.UserBooks, nil
}

func GetUserBookByID(id int) (tablesModels.UserBook, error) {
	if id == 0 {
		return tablesModels.UserBook{}, fmt.Errorf("invalid user book ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/userBooks/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting user book by ID %d: %v", id, err)
		return tablesModels.UserBook{}, fmt.Errorf("error getting user book: %v", err)
	}

	var userBooksResp response.UserBooksResp
	if err := json.Unmarshal(jsonData, &userBooksResp); err != nil {
		log.Printf("Error unmarshalling user book: %v", err)
		return tablesModels.UserBook{}, fmt.Errorf("error unmarshalling user book: %v", err)
	}

	if userBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", userBooksResp.Status.Code)
		return tablesModels.UserBook{}, fmt.Errorf("error getting user book: %s", userBooksResp.Status.Code)
	}

	if len(userBooksResp.UserBooks) == 0 {
		log.Printf("No user book found with ID %d", id)
		return tablesModels.UserBook{}, fmt.Errorf("user book not found")
	}

	return userBooksResp.UserBooks[0], nil
}

func GetUserBooksByUserID(userID int) ([]tablesModels.UserBook, error) {
	if userID == 0 {
		return nil, fmt.Errorf("invalid user ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/userBooks/get/userID/"+strconv.Itoa(userID), nil)
	if err != nil {
		log.Printf("Error getting user books by user ID %d: %v", userID, err)
		return nil, fmt.Errorf("error getting user books: %v", err)
	}

	var userBooksResp response.UserBooksResp
	if err := json.Unmarshal(jsonData, &userBooksResp); err != nil {
		log.Printf("Error unmarshalling user books: %v", err)
		return nil, fmt.Errorf("error unmarshalling user books: %v", err)
	}

	if userBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", userBooksResp.Status.Code)
		return nil, fmt.Errorf("error getting user books: %s", userBooksResp.Status.Code)
	}

	return userBooksResp.UserBooks, nil
}

func GetUserBooksByBookID(bookID int) ([]tablesModels.UserBook, error) {
	if bookID == 0 {
		return nil, fmt.Errorf("invalid book ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/userBooks/get/bookID/"+strconv.Itoa(bookID), nil)
	if err != nil {
		log.Printf("Error getting user books by book ID %d: %v", bookID, err)
		return nil, fmt.Errorf("error getting user books: %v", err)
	}

	var userBooksResp response.UserBooksResp
	if err := json.Unmarshal(jsonData, &userBooksResp); err != nil {
		log.Printf("Error unmarshalling user books: %v", err)
		return nil, fmt.Errorf("error unmarshalling user books: %v", err)
	}

	if userBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", userBooksResp.Status.Code)
		return nil, fmt.Errorf("error getting user books: %s", userBooksResp.Status.Code)
	}

	return userBooksResp.UserBooks, nil
}
