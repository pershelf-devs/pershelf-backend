package userUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateUser(user tablesModels.User) error {
	payload, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshalling user: %v", err)
		return fmt.Errorf("error marshalling user: %w", err)
	}

	jsonData, err := helperContact.HelperRequest("/users/update", payload)
	if err != nil {
		log.Printf("Error making request to update user: %v", err)
		return err
	}

	var userResp response.ResponseMessage
	if err := json.Unmarshal(jsonData, &userResp); err != nil {
		return err
	}

	if userResp.Code != "0" {
		return fmt.Errorf("error updating user: %s", userResp.Code)
	}

	return nil
}
