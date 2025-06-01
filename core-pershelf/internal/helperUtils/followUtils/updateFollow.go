package followUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateFollow(follow tablesModels.Follow) (tablesModels.Follow, error) {
	// Marshal the follow into JSON
	payload, err := json.Marshal(follow)
	if err != nil {
		log.Printf("Error marshalling follow: %v", err)
		return tablesModels.Follow{}, fmt.Errorf("error marshalling follow: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/follows/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.Follow{}, fmt.Errorf("error updating follow: %v", err)
	}

	var followsResp response.FollowsResp
	if err := json.Unmarshal(jsonData, &followsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.Follow{}, fmt.Errorf("error unmarshalling updated follow: %v", err)
	}

	if followsResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", followsResp.Status.Code)
		return tablesModels.Follow{}, fmt.Errorf("error updating follow: %s", followsResp.Status.Code)
	}

	if len(followsResp.Follows) == 0 {
		log.Printf("Follow update succeeded but no follow returned")
		return tablesModels.Follow{}, fmt.Errorf("no follow returned after update")
	}

	return followsResp.Follows[0], nil
}
