package books

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/internal/helperUtils/bookUtils"
	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

func GetBookByIDHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth  = string(ctx.Path())
		book tablesModels.Book
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var bookID int
		if err := json.Unmarshal(ctx.Request.Body(), &bookID); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		if bookID == 0 {
			log.Printf("Invalid book ID at endpoint %s: %d", pth, bookID)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Invalid book ID"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// ✅ Here we call directly to database layer (currently mocked)
		book = bookUtils.GetBookByID(bookID)
		if book.ID == 0 {
			log.Printf("Error getting book by ID at endpoint %s: %d", pth, bookID)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error getting book by ID"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

		// ✅ Send valid JSON response
		if err := json.NewEncoder(ctx).Encode(response.BooksResp{
			Status: response.ResponseMessage{Code: "0"},
			Books:  []tablesModels.Book{book},
		}); err != nil {
			log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error encoding the response body"}}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}
	}
}

func GetBooksByGenreHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var genre string
		if err := json.Unmarshal(ctx.Request.Body(), &genre); err != nil {
			log.Printf("Error unmarshalling genre at endpoint %s: %v", pth, err)
			resp, _ := json.Marshal(response.ResponseMessage{Code: "3", Values: []string{"Invalid genre"}})
			ctx.SetContentType("application/json")
			ctx.SetBody(resp)
			return
		}

		books := bookUtils.GetBooksByGenre(genre)
		if len(books) == 0 {
			log.Printf("No books found for genre at endpoint %s: %s", pth, genre)
			resp, _ := json.Marshal(response.ResponseMessage{Code: "3", Values: []string{"No books found for genre"}})
			ctx.SetContentType("application/json")
			ctx.SetBody(resp)
			return
		}

		resp, _ := json.Marshal(response.BooksResp{
			Status: response.ResponseMessage{Code: "0"},
			Books:  books,
		})
		ctx.SetContentType("application/json")
		ctx.SetBody(resp)
	}
}
