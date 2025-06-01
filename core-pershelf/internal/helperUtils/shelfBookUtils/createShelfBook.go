package shelfBookUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func CreateShelfBook(shelfBook tablesModels.ShelfBook) (tablesModels.ShelfBook, error) {
	// Marshal the shelfBook into JSON
	payload, err := json.Marshal(shelfBook)
	if err != nil {
		log.Printf("Error marshalling shelf book: %v", err)
		return tablesModels.ShelfBook{}, fmt.Errorf("error marshalling shelf book: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/shelfBooks/create", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.ShelfBook{}, fmt.Errorf("error creating shelf book: %v", err)
	}

	var shelfBooksResp response.ShelfBooksResp
	if err := json.Unmarshal(jsonData, &shelfBooksResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.ShelfBook{}, fmt.Errorf("error unmarshalling created shelf book: %v", err)
	}

	if shelfBooksResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error with code: : %s", shelfBooksResp.Status.Code)
		return tablesModels.ShelfBook{}, fmt.Errorf("error creating shelf book: %s", shelfBooksResp.Status.Code)
	}

	if len(shelfBooksResp.ShelfBooks) == 0 {
		log.Printf("Shelf book creation succeeded but no shelf book returned")
		return tablesModels.ShelfBook{}, fmt.Errorf("no shelf book returned after creation")
	}

	return shelfBooksResp.ShelfBooks[0], nil
}
