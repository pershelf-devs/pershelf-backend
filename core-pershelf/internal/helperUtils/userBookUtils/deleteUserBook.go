package userBookUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteUserBook(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid user book ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/userBooks/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting user book: %v", err)
	}

	var userBooksResp response.UserBooksResp
	if err := json.Unmarshal(jsonData, &userBooksResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if userBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", userBooksResp.Status.Code)
		return fmt.Errorf("error deleting user book: %s", userBooksResp.Status.Code)
	}

	return nil
}
