package bookLikes

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/core-pershelf/rest/helperContact/tablesModels"
	"github.com/valyala/fasthttp"
)

func CreateBookLikeHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var bookLike tablesModels.BookLike
		if err := json.Unmarshal(ctx.Request.Body(), &bookLike); err != nil {
			log.Printf("Error unmarshalling the request body at endpoint %s: %v", pth, err)
			log.Printf("Request body: %s", string(ctx.Request.Body()))
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3", Values: []string{"Error unmarshalling request body"}}); err != nil {
				log.Printf("(Error): error encoding response message at endpoint (%s).", pth)
			}
			return
		}

	}
}
