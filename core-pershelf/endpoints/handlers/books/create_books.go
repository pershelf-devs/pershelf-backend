package books

import (
	"context"
	"encoding/json"
	"log"

	"github.com/core-pershelf/globals"
	"github.com/core-pershelf/mongo/tablesModels"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

// CreateBookHandler creates a new book
func CreateBookHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var book tablesModels.Book
		if err := json.Unmarshal(ctx.Request.Body(), &book); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "3",
				Values: []string{"Invalid request format"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Check if the required fields are provided
		if book.Title == "" || book.Author == "" || book.ISBN == "" || book.Category == "" || book.Language == "" {
			log.Printf("Missing required fields at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "4",
				Values: []string{"All required fields must be provided"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Check if the page count is valid
		if book.PageCount <= 0 {
			log.Printf("Invalid page count at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "6",
				Values: []string{"Invalid page count"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Check if the language code is valid
		if len(book.Language) != 2 && len(book.Language) != 3 {
			log.Printf("Invalid language code at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "7",
				Values: []string{"Invalid language code"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Save the book to the database
		if _, err := globals.BooksCollection.InsertOne(context.Background(), book); err != nil {
			log.Printf("Error creating book at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "9",
				Values: []string{"Error creating book"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Return the success message
		if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
			Code:   "0",
			Values: []string{"Book created successfully"},
		}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
	}
}
