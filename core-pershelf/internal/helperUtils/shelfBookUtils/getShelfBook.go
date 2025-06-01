package shelfBookUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllShelfBooks() ([]tablesModels.ShelfBook, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/shelfBooks/get/all", nil)
	if err != nil {
		log.Printf("Error getting all shelf books: %v", err)
		return nil, fmt.Errorf("error getting all shelf books: %v", err)
	}

	var shelfBooksResp response.ShelfBooksResp
	if err := json.Unmarshal(jsonData, &shelfBooksResp); err != nil {
		log.Printf("Error unmarshalling shelf books: %v", err)
		return nil, fmt.Errorf("error unmarshalling shelf books: %v", err)
	}

	if shelfBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", shelfBooksResp.Status.Code)
		return nil, fmt.Errorf("error getting shelf books: %s", shelfBooksResp.Status.Code)
	}

	return shelfBooksResp.ShelfBooks, nil
}

func GetShelfBookByID(id int) (tablesModels.ShelfBook, error) {
	if id == 0 {
		return tablesModels.ShelfBook{}, fmt.Errorf("invalid shelf book ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/shelfBooks/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting shelf book by ID %d: %v", id, err)
		return tablesModels.ShelfBook{}, fmt.Errorf("error getting shelf book: %v", err)
	}

	var shelfBooksResp response.ShelfBooksResp
	if err := json.Unmarshal(jsonData, &shelfBooksResp); err != nil {
		log.Printf("Error unmarshalling shelf book: %v", err)
		return tablesModels.ShelfBook{}, fmt.Errorf("error unmarshalling shelf book: %v", err)
	}

	if shelfBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", shelfBooksResp.Status.Code)
		return tablesModels.ShelfBook{}, fmt.Errorf("error getting shelf book: %s", shelfBooksResp.Status.Code)
	}

	if len(shelfBooksResp.ShelfBooks) == 0 {
		log.Printf("No shelf book found with ID %d", id)
		return tablesModels.ShelfBook{}, fmt.Errorf("shelf book not found")
	}

	return shelfBooksResp.ShelfBooks[0], nil
}
