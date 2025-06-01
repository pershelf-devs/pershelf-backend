package userShelfUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllUserShelfs() ([]tablesModels.UserShelf, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/userShelfs/get/all", nil)
	if err != nil {
		log.Printf("Error getting all user shelfs: %v", err)
		return nil, fmt.Errorf("error getting all user shelfs: %v", err)
	}

	var userShelfsResp response.UserShelfsResp
	if err := json.Unmarshal(jsonData, &userShelfsResp); err != nil {
		log.Printf("Error unmarshalling user shelfs: %v", err)
		return nil, fmt.Errorf("error unmarshalling user shelfs: %v", err)
	}

	if userShelfsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", userShelfsResp.Status.Code)
		return nil, fmt.Errorf("error getting user shelfs: %s", userShelfsResp.Status.Code)
	}

	return userShelfsResp.UserShelfs, nil
}

func GetUserShelfByID(id int) (tablesModels.UserShelf, error) {
	if id == 0 {
		return tablesModels.UserShelf{}, fmt.Errorf("invalid user shelf ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/userShelfs/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting user shelf by ID %d: %v", id, err)
		return tablesModels.UserShelf{}, fmt.Errorf("error getting user shelf: %v", err)
	}

	var userShelfsResp response.UserShelfsResp
	if err := json.Unmarshal(jsonData, &userShelfsResp); err != nil {
		log.Printf("Error unmarshalling user shelf: %v", err)
		return tablesModels.UserShelf{}, fmt.Errorf("error unmarshalling user shelf: %v", err)
	}

	if userShelfsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", userShelfsResp.Status.Code)
		return tablesModels.UserShelf{}, fmt.Errorf("error getting user shelf: %s", userShelfsResp.Status.Code)
	}

	if len(userShelfsResp.UserShelfs) == 0 {
		log.Printf("No user shelf found with ID %d", id)
		return tablesModels.UserShelf{}, fmt.Errorf("user shelf not found")
	}

	return userShelfsResp.UserShelfs[0], nil
}
