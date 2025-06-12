package userBookRelationUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func UpdateUserBookRelation(userBookRelation tablesModels.UserBookRelation) (string, error) {
	// Marshal the userBookRelation into JSON
	payload, err := json.Marshal(userBookRelation)
	if err != nil {
		log.Printf("Error marshalling user book relation: %v", err)
		return "3", fmt.Errorf("error marshalling user book relation: %v", err)
	}

	jsonData, err := helperContact.HelperRequest("/user-book-relations/update", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return "3", fmt.Errorf("error calling helper request: %v", err)
	}

	var userBookRelationResponse response.UserBookRelationsResp
	if err := json.Unmarshal(jsonData, &userBookRelationResponse); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return "3", fmt.Errorf("error unmarshalling response: %v", err)
	}

	if len(userBookRelationResponse.UserBookRelations) == 0 {
		log.Printf("User book relation not updated")
		return "3", fmt.Errorf("user book relation not updated")
	}

	if userBookRelationResponse.Status.Code != "0" {
		log.Printf("Error updating user book relation: %v", userBookRelationResponse.Status.Values)
		return "3", fmt.Errorf("error updating user book relation: %v", userBookRelationResponse.Status.Values)
	}

	return "0", nil
}
