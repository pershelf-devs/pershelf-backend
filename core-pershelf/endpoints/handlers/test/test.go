package test

import (
	"encoding/json"
	"log"

	"github.com/valyala/fasthttp"
)

func ExecuteTestHandler(ctx *fasthttp.RequestCtx) {
	var (
		pth = string(ctx.Path())
	)

	select {
	case <-ctx.Done():
		log.Printf("Client canceled the request at endpoint (%s).", pth)
		ctx.SetBody(nil)
		return

	default:
		// Prepare a test struct
		type TestApiStruct struct {
			Status  string `json:"status"`
			Message string `json:"message"`
		}

		testApi := TestApiStruct{
			Status:  "200",
			Message: "Api is working successfully",
		}

		// Send message in a response to client's request.
		if err := json.NewEncoder(ctx).Encode(testApi); err != nil {
			log.Printf("Error encoding response at endpoint (%s): %v", pth, err)
		}
	}
}
