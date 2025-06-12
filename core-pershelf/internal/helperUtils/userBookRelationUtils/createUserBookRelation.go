package userBookRelationUtils

import (
	"encoding/json"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func CreateUserBookRelation(userBookRelation tablesModels.UserBookRelation) tablesModels.UserBookRelation {
	// Marshal the userBookRelation into JSON
	payload, err := json.Marshal(userBookRelation)
	if err != nil {
		log.Printf("Error marshalling user book relation: %v", err)
		return tablesModels.UserBookRelation{}
	}

	jsonData, err := helperContact.HelperRequest("/user-book-relations/create", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.UserBookRelation{}
	}

	var userBookRelationResponse response.UserBookRelationsResp
	if err := json.Unmarshal(jsonData, &userBookRelationResponse); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.UserBookRelation{}
	}

	if len(userBookRelationResponse.UserBookRelations) == 0 {
		log.Printf("User book relation not created")
		return tablesModels.UserBookRelation{}
	}

	if userBookRelationResponse.Status.Code != "0" {
		log.Printf("Error creating user book relation: %v", userBookRelationResponse.Status.Values)
		return tablesModels.UserBookRelation{}
	}

	return userBookRelationResponse.UserBookRelations[0]
}
