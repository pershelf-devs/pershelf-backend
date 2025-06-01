package userUtils

import (
	"encoding/json"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

// GetAllUsers gets all users from the database
func GetAllUsers() []tablesModels.User {
	jsonData, err := helperContact.HelperRequest("/users/get/all", nil)
	if err != nil {
		log.Printf("Error getting all users: %v", err)
		return nil
	}

	var userResp response.UsersResp
	if err := json.Unmarshal(jsonData, &userResp); err != nil {
		log.Printf("Error unmarshalling user: %v", err)
		return nil
	}

	if userResp.Status.Code != "0" {
		log.Printf("Error getting all users: %v", userResp.Status.Code)
		return nil
	}

	return userResp.Users
}

// GetUserByID gets a user by id from the database
func GetUserByID(id int) tablesModels.User {
	if id == 0 {
		return tablesModels.User{}
	}

	jsonData, err := helperContact.HelperRequest("/users/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting user by id: %v", err)
		return tablesModels.User{}
	}

	var userResp response.UsersResp
	if err := json.Unmarshal(jsonData, &userResp); err != nil {
		log.Printf("Error unmarshalling user: %v", err)
		return tablesModels.User{}
	}

	if len(userResp.Users) == 0 {
		log.Printf("User with id (%d) not found", id)
		return tablesModels.User{}
	}

	return userResp.Users[0]
}

// GetUserByEmail gets a user by email from the database
func GetUserByEmail(email string) tablesModels.User {
	if email == "" {
		log.Println("Email is empty, cannot get user by email")
		return tablesModels.User{}
	}

	payload, err := json.Marshal(email)
	if err != nil {
		log.Printf("Error marshalling email: %v", err)
		return tablesModels.User{}
	}

	jsonData, err := helperContact.HelperRequest("/users/get/by-email", payload)
	if err != nil {
		log.Printf("Error getting user by email: %v", err)
		return tablesModels.User{}
	}

	var userResp response.UsersResp
	if err := json.Unmarshal(jsonData, &userResp); err != nil {
		log.Printf("Error unmarshalling user: %v", err)
		return tablesModels.User{}
	}

	if len(userResp.Users) == 0 {
		log.Printf("User with email (%s) not found", email)
		return tablesModels.User{}
	}

	return userResp.Users[0]
}
