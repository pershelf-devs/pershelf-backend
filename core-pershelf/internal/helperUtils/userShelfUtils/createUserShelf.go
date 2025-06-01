package userShelfUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func CreateUserShelf(userShelf tablesModels.UserShelf) (tablesModels.UserShelf, error) {
	// Marshal the userShelf into JSON
	payload, err := json.Marshal(userShelf)
	if err != nil {
		log.Printf("Error marshalling user shelf: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error marshalling user shelf: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/userShelfs/create", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error creating user shelf: %v", err)
	}

	var userShelfsResp response.UserShelfsResp
	if err := json.Unmarshal(jsonData, &userShelfsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error unmarshalling created user shelf: %v", err)
	}

	if userShelfsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", userShelfsResp.Status.Code)
		return tablesModels.UserShelf{}, fmt.Errorf("error creating user shelf: %s", userShelfsResp.Status.Code)
	}

	if len(userShelfsResp.UserShelfs) == 0 {
		log.Printf("User shelf creation succeeded but no user shelf returned")
		return tablesModels.UserShelf{}, fmt.Errorf("no user shelf returned after creation")
	}

	return userShelfsResp.UserShelfs[0], nil
}
