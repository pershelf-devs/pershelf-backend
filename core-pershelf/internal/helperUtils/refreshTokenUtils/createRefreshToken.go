package refreshTokenUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func CreateRefreshToken(refreshToken tablesModels.RefreshToken) (tablesModels.RefreshToken, error) {
	// Marshal the refresh token into JSON
	payload, err := json.Marshal(refreshToken)
	if err != nil {
		log.Printf("Error marshalling refresh token: %v", err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error marshalling refresh token: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/refreshTokens/create", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error creating refresh token: %v", err)
	}

	var tokenResp response.RefreshTokensResp
	if err := json.Unmarshal(jsonData, &tokenResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.RefreshToken{}, fmt.Errorf("error unmarshalling created refresh token: %v", err)
	}

	if tokenResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", tokenResp.Status.Code)
		return tablesModels.RefreshToken{}, fmt.Errorf("error creating refresh token: %s", tokenResp.Status.Code)
	}

	if len(tokenResp.RefreshTokens) == 0 {
		log.Printf("Refresh token creation succeeded but no token returned")
		return tablesModels.RefreshToken{}, fmt.Errorf("no refresh token returned after creation")
	}

	return tokenResp.RefreshTokens[0], nil
}
