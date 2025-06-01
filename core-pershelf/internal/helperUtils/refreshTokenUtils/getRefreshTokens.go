package refreshTokenUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func GetAllRefreshTokens() ([]tablesModels.RefreshToken, error) {
	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/refreshTokens/get/all", nil)
	if err != nil {
		log.Printf("Error getting all refresh tokens: %v", err)
		return nil, fmt.Errorf("error getting all refresh tokens: %v", err)
	}

	var tokenResp response.RefreshTokensResp
	if err := json.Unmarshal(jsonData, &tokenResp); err != nil {
		log.Printf("Error unmarshalling refresh tokens: %v", err)
		return nil, fmt.Errorf("error unmarshalling refresh tokens: %v", err)
	}

	if tokenResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", tokenResp.Status.Code)
		return nil, fmt.Errorf("error getting refresh tokens: %s", tokenResp.Status.Code)
	}

	return tokenResp.RefreshTokens, nil
}

func GetRefreshTokenByID(id int) (tablesModels.RefreshToken, error) {
	if id == 0 {
		return tablesModels.RefreshToken{}, fmt.Errorf("invalid refresh token ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/refreshTokens/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting refresh token by ID %d: %v", id, err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error getting refresh token: %v", err)
	}

	var tokenResp response.RefreshTokensResp
	if err := json.Unmarshal(jsonData, &tokenResp); err != nil {
		log.Printf("Error unmarshalling refresh token: %v", err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error unmarshalling refresh token: %v", err)
	}

	if tokenResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", tokenResp.Status.Code)
		return tablesModels.RefreshToken{}, fmt.Errorf("error getting refresh token: %s", tokenResp.Status.Code)
	}

	if len(tokenResp.RefreshTokens) == 0 {
		log.Printf("No refresh token found with ID %d", id)
		return tablesModels.RefreshToken{}, fmt.Errorf("refresh token not found")
	}

	return tokenResp.RefreshTokens[0], nil
}

func GetRefreshTokenByUserID(userID int) (tablesModels.RefreshToken, error) {
	if userID == 0 {
		return tablesModels.RefreshToken{}, fmt.Errorf("invalid user ID")
	}

	// Call the helper request
	jsonData, err := helperContact.HelperRequest("/refreshTokens/get/userID/"+strconv.Itoa(userID), nil)
	if err != nil {
		log.Printf("Error getting refresh token by user ID %d: %v", userID, err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error getting refresh token: %v", err)
	}

	var tokenResp response.RefreshTokensResp
	if err := json.Unmarshal(jsonData, &tokenResp); err != nil {
		log.Printf("Error unmarshalling refresh token: %v", err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error unmarshalling refresh token: %v", err)
	}

	if tokenResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", tokenResp.Status.Code)
		return tablesModels.RefreshToken{}, fmt.Errorf("error getting refresh token: %s", tokenResp.Status.Code)
	}

	if len(tokenResp.RefreshTokens) == 0 {
		log.Printf("No refresh token found for user ID %d", userID)
		return tablesModels.RefreshToken{}, fmt.Errorf("refresh token not found")
	}

	return tokenResp.RefreshTokens[0], nil
}
