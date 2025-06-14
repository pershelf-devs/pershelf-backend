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

func GetBooksByIDs(ids []int) []tablesModels.Book {
	if len(ids) == 0 {
		log.Printf("No book IDs provided")
		return nil
	}

	// Marshal ids to json
	payload, err := json.Marshal(ids)
	if err != nil {
		log.Printf("Error marshalling ids: %v", err)
		return nil
	}

	jsonData, err := helperContact.HelperRequest("/books/get/ids", payload)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return nil
	}

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling book: %v", err)
		return nil
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Error getting books by IDs: %v", bookResp.Status.Code)
		return nil
	}

	return bookResp.Books
}

func GetBookByID(id int) tablesModels.Book {
	if id == 0 {
		log.Printf("Invalid book ID")
		return tablesModels.Book{}
	}

	jsonData, err := helperContact.HelperRequest("/books/get/id/"+strconv.Itoa(id), nil)
	if err != nil {
		log.Printf("Error calling helper request: %v", err)
		return tablesModels.Book{}
	}

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling book: %v", err)
		return tablesModels.Book{}
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Backend returned error: %s", bookResp.Status.Code)
		return tablesModels.Book{}
	}

	if len(bookResp.Books) == 0 {
		log.Printf("No book found with ID %d", id)
		return tablesModels.Book{}
	}

	return bookResp.Books[0]
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

// GetBooksByGenre retrieves books by genre from the database.
func GetBooksByGenre(genre string) []tablesModels.Book {
	if genre == "" {
		log.Printf("Genre is empty")
		return nil
	}

	jsonData, err := helperContact.HelperRequest("/books/get/genre/"+genre, nil)
	if err != nil {
		log.Printf("Error getting books by genre %s: %v", genre, err)
		return nil
	}

	var bookResp response.BooksResp
	if err := json.Unmarshal(jsonData, &bookResp); err != nil {
		log.Printf("Error unmarshalling books: %v", err)
		return nil
	}

	if bookResp.Status.Code != "0" {
		log.Printf("Backend returned error for genre %s: %s", genre, bookResp.Status.Code)
		return nil
	}

	if len(bookResp.Books) == 0 {
		log.Printf("No books found for genre %s", genre)
		return nil
	}

	return bookResp.Books
}
