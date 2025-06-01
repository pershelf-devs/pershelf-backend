package userUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteUserByID(id int) error {
	payload := map[string]int{"id": id}
	jsonData, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling payload: %v", err)
		return fmt.Errorf("error marshalling payload: %w", err)
	}

	responseData, err := helperContact.HelperRequest(fmt.Sprintf("/users/delete/id/%d", id), jsonData)
	if err != nil {
		log.Printf("Error making request to delete user: %v", err)
		return err
	}

	var userResp response.UsersResp
	if err := json.Unmarshal(responseData, &userResp); err != nil {
		return err
	}

	if userResp.Status.Code != "0" {
		return fmt.Errorf("error deleting user: %s", userResp.Status.Code)
	}

	return nil
}
