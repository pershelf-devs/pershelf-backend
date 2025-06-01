package bookUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateBook(book tablesModels.Book) error {
	// Marshal the book into JSON
	payload, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error marshalling book: %v", err)
		return fmt.Errorf("error marshalling book: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/books/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error updating book: %v", err)
	}

	var updateResp response.ResponseMessage
	if err := json.Unmarshal(jsonData, &updateResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling updated book: %v", err)
	}

	if updateResp.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", updateResp.Code)
		return fmt.Errorf("error updating book: %s", updateResp.Code)
	}

	return nil
}
