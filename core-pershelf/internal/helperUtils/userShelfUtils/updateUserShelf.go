package userShelfUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateUserShelf(userShelf tablesModels.UserShelf) (tablesModels.UserShelf, error) {
	// Marshal the userShelf into JSON
	payload, err := json.Marshal(userShelf)
	if err != nil {
		log.Printf("Error marshalling user shelf: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error marshalling user shelf: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/userShelfs/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error updating user shelf: %v", err)
	}

	var userShelfsResp response.UserShelfsResp
	if err := json.Unmarshal(jsonData, &userShelfsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error unmarshalling updated user shelf: %v", err)
	}

	if userShelfsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", userShelfsResp.Status.Code)
		return tablesModels.UserShelf{}, fmt.Errorf("error updating user shelf: %s", userShelfsResp.Status.Code)
	}

	if len(userShelfsResp.UserShelfs) == 0 {
		log.Printf("User shelf update succeeded but no user shelf returned")
		return tablesModels.UserShelf{}, fmt.Errorf("no user shelf returned after update")
	}

	return userShelfsResp.UserShelfs[0], nil
}
