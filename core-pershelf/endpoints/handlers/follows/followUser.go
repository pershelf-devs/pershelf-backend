package follows

import (
	"encoding/json"
	"log"

	"github.com/core-pershelf/rest/helperContact/response"
	"github.com/valyala/fasthttp"
)

func FollowUserHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		var userID int
		if err := json.Unmarshal(ctx.Request.Body(), &userID); err != nil {
			log.Printf("Error unmarshling request body at endpoint (%s)", pth)
			if err := json.NewEncoder(ctx).Encode(response.ResponseMessage{Code: "3"}); err != nil {
				log.Printf("Error encoding the response body at endpoint %s: %v", pth, err)
			}
			return
		}

	}
}
