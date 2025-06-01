package bookUtils

import (
	"encoding/json"
	"fmt"
	"log"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

func CreateBook(book tablesModels.Book) (tablesModels.Book, error) {
	// Marshal the book into JSON
	payload, err := json.Marshal(book)
	if err != nil {
		log.Printf("Error marshalling book: %v", err)
		return tablesModels.Book{}, fmt.Errorf("error marshalling book: %w", err)
	}

	// Call the helper request with marshalled JSON
	jsonData, err := helperContact.HelperRequest("/books/create", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.Book{}, fmt.Errorf("error creating book: %v", err)
	}

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling response: %v", err)
		return tablesModels.Book{}, fmt.Errorf("error unmarshalling created book: %v", err)
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Helper microservice returned an error: %+v", bookResp.Status)
		return tablesModels.Book{}, fmt.Errorf("error creating book: %s", bookResp.Status.Code)
	}

	if len(bookResp.Books) == 0 {
		log.Printf("Book creation succeeded but no book returned")
		return tablesModels.Book{}, fmt.Errorf("no book returned after creation")
	}

	return bookResp.Books[0], nil
}
