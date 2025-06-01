package userUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateUser(user tablesModels.User) (tablesModels.User, error) {
	payload, err := json.Marshal(user)
	if err != nil {
		log.Printf("Error marshalling user: %v", err)
		return tablesModels.User{}, fmt.Errorf("error marshalling user: %w", err)
	}

	jsonData, err := helperContact.HelperRequest("/users/update", payload)
	if err != nil {
		log.Printf("Error making request to update user: %v", err)
		return tablesModels.User{}, err
	}

	var userResp response.UsersResp
	if err := json.Unmarshal(jsonData, &userResp); err != nil {
		return tablesModels.User{}, err
	}

	if userResp.Status.Code != "0" {
		return tablesModels.User{}, fmt.Errorf("error updating user: %s", userResp.Status.Code)
	}

	if len(userResp.Users) == 0 {
		return tablesModels.User{}, fmt.Errorf("no user updated")
	}

	return userResp.Users[0], nil
}
