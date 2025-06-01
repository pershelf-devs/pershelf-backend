package userShelfUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteUserShelf(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid user shelf ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/userShelfs/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting user shelf: %v", err)
	}

	var userShelfsResp response.UserShelfsResp
	if err := json.Unmarshal(jsonData, &userShelfsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if userShelfsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", userShelfsResp.Status.Code)
		return fmt.Errorf("error deleting user shelf: %s", userShelfsResp.Status.Code)
	}

	return nil
}
