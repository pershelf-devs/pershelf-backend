package refreshTokenUtils

import (
	"encoding/json"
	"fmt"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
)

func DeleteRefreshToken(id int) error {
	if id == 0 {
		return fmt.Errorf("invalid refresh token ID")
	}

	// Call the helper request (no payload needed)
	jsonData, err := helperContact.HelperRequest("/refreshTokens/delete/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return fmt.Errorf("error deleting refresh token: %v", err)
	}

	var tokenResp response.RefreshTokensResp
	if err := json.Unmarshal(jsonData, &tokenResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return fmt.Errorf("error unmarshalling delete response: %v", err)
	}

	if tokenResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", tokenResp.Status.Code)
		return fmt.Errorf("error deleting refresh token: %s", tokenResp.Status.Code)
	}

	return nil
}
