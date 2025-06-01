package bookUtils

import (
	"encoding/json"
	"log"
	"strconv"

	helperContact "github.com/core-pershelf/rest/helperContact/request"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
)

// GetAllBooks gets all books from the database
func GetAllBooks() []tablesModels.Book {
	jsonData, err := helperContact.HelperRequest("/books/get/all", nil)
	if err != nil {
		log.Printf("Error getting all books: %v", err)
		return nil
	}

	// Debug log
	log.Printf("All books json data: %s", string(jsonData))

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling books: %v", err)
		return nil
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Error getting all books: %v", bookResp.Status.Code)
		return nil
	}

	return bookResp.Books
}

func GetBookByID(id int) *tablesModels.Book {
	if id == 0 {
		return nil
	}

	jsonData, err := helperContact.HelperRequest("/books/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error getting book by ID %d: %v", id, err)
		return nil
	}

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling book: %v", err)
		return nil
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Error getting book by ID %d: %v", id, bookResp.Status.Code)
		return nil
	}

	if len(bookResp.Books) == 0 {
		log.Printf("No book found with ID %d", id)
		return nil
	}

	return &bookResp.Books[0]
}

func GetBookByISBN(isbn string) *tablesModels.Book {
	if isbn == "" {
		return nil
	}

	jsonData, err := helperContact.HelperRequest("/books/get/isbn/"+isbn, nil)
	if err != nil {
		log.Printf("Error getting book by ISBN %s: %v", isbn, err)
		return nil
	}

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling book: %v", err)
		return nil
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Error getting book by ISBN %s: %v", isbn, bookResp.Status.Code)
		return nil
	}

	if len(bookResp.Books) == 0 {
		log.Printf("No book found with ISBN %s", isbn)
		return nil
	}

	return &bookResp.Books[0]
}
