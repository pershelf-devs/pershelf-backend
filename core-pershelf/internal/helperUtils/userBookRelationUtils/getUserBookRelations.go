package userBookRelationUtils

import (
	"encoding/json"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

// GetUserBookRelationsByUserID gets all user book relations by user id
func GetUserBookRelationsByUserID(userID int) []tablesModels.UserBookRelation {
	jsonData, err := helperContact.HelperRequest("/user-book-relations/get/user-id/"+strconv.Itoa(userID), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return []tablesModels.UserBookRelation{}
	}

	var userBookRelationsResp response.UserBookRelationsResp
	if err := json.Unmarshal(jsonData, &userBookRelationsResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return []tablesModels.UserBookRelation{}
	}

	if userBookRelationsResp.Status.Code != "0" {
		log.Printf("Error getting user book relations: %v", userBookRelationsResp.Status.Values)
		return []tablesModels.UserBookRelation{}
	}

	if len(userBookRelationsResp.UserBookRelations) == 0 {
		log.Printf("No user book relations found for user ID: %d", userID)
		return []tablesModels.UserBookRelation{}
	}

	return userBookRelationsResp.UserBookRelations
}

func GetUserBookRelationByUserIDAndBookID(userID, bookID int) tablesModels.UserBookRelation {
	jsonData, err := helperContact.HelperRequest("/user-book-relations/get/user-id/"+strconv.Itoa(userID)+"/book-id/"+strconv.Itoa(bookID), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.UserBookRelation{}
	}

	var userBookRelation response.UserBookRelationsResp
	if err := json.Unmarshal(jsonData, &userBookRelation); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.UserBookRelation{}
	}

	if len(userBookRelation.UserBookRelations) == 0 {
		log.Printf("User book relation not found for user ID: %d and book ID: %d", userID, bookID)
		return tablesModels.UserBookRelation{}
	}

	return userBookRelation.UserBookRelations[0]
}
