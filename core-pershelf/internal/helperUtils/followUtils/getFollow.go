package followUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllFollows() ([]tablesModels.Follow, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/follows/get/all", nil)
	if err != nil {
		log.Printf("Error getting all follows: %v", err)
		return nil, fmt.Errorf("error getting all follows: %v", err)
	}

	var followsResp response.FollowsResp
	if err := json.Unmarshal(jsonData, &followsResp); err != nil {
		log.Printf("Error unmarshalling follows: %v", err)
		return nil, fmt.Errorf("error unmarshalling follows: %v", err)
	}

	if followsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", followsResp.Status.Code)
		return nil, fmt.Errorf("error getting follows: %s", followsResp.Status.Code)
	}

	return followsResp.Follows, nil
}

func GetFollowByID(id int) (tablesModels.Follow, error) {
	if id == 0 {
		return tablesModels.Follow{}, fmt.Errorf("invalid follow ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/follows/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting follow by ID %d: %v", id, err)
		return tablesModels.Follow{}, fmt.Errorf("error getting follow: %v", err)
	}

	var followsResp response.FollowsResp
	if err := json.Unmarshal(jsonData, &followsResp); err != nil {
		log.Printf("Error unmarshalling follow: %v", err)
		return tablesModels.Follow{}, fmt.Errorf("error unmarshalling follow: %v", err)
	}

	if followsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", followsResp.Status.Code)
		return tablesModels.Follow{}, fmt.Errorf("error getting follow: %s", followsResp.Status.Code)
	}

	if len(followsResp.Follows) == 0 {
		log.Printf("No follow found with ID %d", id)
		return tablesModels.Follow{}, fmt.Errorf("follow not found")
	}

	return followsResp.Follows[0], nil
}
