package followUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteFollow(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid follow ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/follows/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting follow: %v", err)
	}

	var followsResp response.FollowsResp
	if err := json.Unmarshal(jsonData, &followsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if followsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", followsResp.Status.Code)
		return fmt.Errorf("error deleting follow: %s", followsResp.Status.Code)
	}

	return nil
}
