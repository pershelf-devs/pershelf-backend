package books

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/bookUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

// GetMostReadBooksHandler returns the most read books (default limit is 10)
func GetMostReadBooksHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth   = string(ctx.Path())
		books []tablesModels.Book
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var request struct {
			Limit    int    `json:"limit"`
			Category string `json:"category"`
		}

		if err := json.Unmarshal(ctx.Request.Body(), &request); err != nil {
			request.Limit = 10
		}

		// Get the most read books
		if books = bookUtils.GetAllBooks(); books == nil {
			log.Printf("Error getting the most read books")
			if err := json.NewEncoder(ctx).Encode(response.BooksResp{
				Status: response.ResponseMessage{Code: "3", Values: []string{"Error getting the most read books"}},
				Books:  nil,
			}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// Limit the number of books
		if request.Limit > 0 {
			if request.Limit > len(books) {
				request.Limit = len(books)
			}
			books = books[:request.Limit]
		}

		// Return the results
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "0", Values: nil},
			Books:  books,
		}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
		}
	}
}
