package shelfBookUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteShelfBook(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid shelf book ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/shelfBooks/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting shelf book: %v", err)
	}

	var shelfBooksResp response.ShelfBooksResp
	if err := json.Unmarshal(jsonData, &shelfBooksResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if shelfBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", shelfBooksResp.Status.Code)
		return fmt.Errorf("error deleting shelf book: %s", shelfBooksResp.Status.Code)
	}

	return nil
}

// DeleteShelfBookByShelfIDAndBookID deletes a shelf_book entry from the database by shelf ID and book ID
func DeleteShelfBookByShelfIDAndBookID(shelfID int, bookID int) error {
	if shelfID == 0 {
		return fmt.Errorf("invalid shelf ID")
	}

	if bookID == 0 {
		return fmt.Errorf("invalid book ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/shelf-books/delete/shelf-id/"+strconv.Itoa(shelfID)+"/book-id/"+strconv.Itoa(bookID), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting shelf book: %v", err)
	}

	var shelfBooksResp response.ResponseMessage
	if err := json.Unmarshal(jsonData, &shelfBooksResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if shelfBooksResp.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", shelfBooksResp.Code)
		return fmt.Errorf("error deleting shelf book: %s, %v", shelfBooksResp.Code, shelfBooksResp.Values)
	}

	return nil
}
