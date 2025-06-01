package userBookUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateUserBook(userBook tablesModels.UserBook) (tablesModels.UserBook, error) {
	// Marshal the userBook into JSON
	payload, err := json.Marshal(userBook)
	if err != nil {
		log.Printf("Error marshalling user book: %v", err)
		return tablesModels.UserBook{}, fmt.Errorf("error marshalling user book: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/userBooks/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.UserBook{}, fmt.Errorf("error updating user book: %v", err)
	}

	var userBooksResp response.UserBooksResp
	if err := json.Unmarshal(jsonData, &userBooksResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.UserBook{}, fmt.Errorf("error unmarshalling updated user book: %v", err)
	}

	if userBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", userBooksResp.Status.Code)
		return tablesModels.UserBook{}, fmt.Errorf("error updating user book: %s", userBooksResp.Status.Code)
	}

	if len(userBooksResp.UserBooks) == 0 {
		log.Printf("User book update succeeded but no user book returned")
		return tablesModels.UserBook{}, fmt.Errorf("no user book returned after update")
	}

	return userBooksResp.UserBooks[0], nil
}
