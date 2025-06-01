package bookUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteBook(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid book ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/books/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting book: %v", err)
	}

	var deleteResp response.ResponseMessage
	if err := json.Unmarshal(jsonData, &deleteResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if deleteResp.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", deleteResp.Code)
		return fmt.Errorf("error deleting book: %s", deleteResp.Code)
	}

	return nil
}
