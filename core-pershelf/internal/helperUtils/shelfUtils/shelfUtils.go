package shelfUtils

import (
	"encoding/json"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

// GetShelfByUserIDAndName returns the shelf for the given user and shelf name.
func GetShelfByUserIDAndName(userID int, shelfName string) tablesModels.UserShelf {
	jsonData, err := helperContact.HelperRequest("/user-shelfs/get/user-id/"+strconv.Itoa(userID)+"/name/"+shelfName, nil)
	if err != nil {
		log.Printf("Error getting shelves for user %d: %v", userID, err)
		return tablesModels.UserShelf{}
	}

	var shelvesResp response.UserShelfsResp
	if err := json.Unmarshal(jsonData, &shelvesResp); err != nil {
		log.Printf("Error unmarshalling shelves: %v", err)
		return tablesModels.UserShelf{}
	}

	if len(shelvesResp.UserShelfs) == 0 {
		log.Printf("Error getting shelves for user %d: %v", userID, err)
		return tablesModels.UserShelf{}
	}

	return shelvesResp.UserShelfs[0]
}
