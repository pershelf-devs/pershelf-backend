package books

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/bookUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
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
		if book.Title == "" || book.Author == "" || book.ISBN == "" {
			log.Printf("Missing required fields at endpoint %s", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "4",
				Values: []string{"All required fields must be provided"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Create the book
		book, err := bookUtils.CreateBook(book)
		if err != nil {
			log.Printf("Error creating book at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "3",
				Values: []string{"Error creating book"},
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		if book.ID == 0 {
			log.Printf("Error creating book at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{
				Code:   "3",
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
