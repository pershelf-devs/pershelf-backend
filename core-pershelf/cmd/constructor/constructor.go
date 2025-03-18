package constructor

import (
	"log"

	"github.com/core-pershelf/endpoints/handlers/auth"
	"github.com/core-pershelf/endpoints/handlers/test"
	"github.com/valyala/fasthttp"
)

func ConstructServer() *fasthttp.Server {
	srv := &fasthttp.Server{
		Handler: MainHandler,
	}

	// Very important to be able to read big-sized requests' headers and bodies. >> default is : 4096
	srv.ReadBufferSize = 10000
	// Very important to be able to write big-sized responses. >> default is : 4096
	srv.WriteBufferSize = 10000

	return srv
}

// MainHandler is the main handler for the fasthttp server.
func MainHandler(ctx *fasthttp.RequestCtx) {

	switch pth := string(ctx.Path()); pth {

	case "/test":
		test.ExecuteTestHandler(ctx)

	case "/login":
		auth.ClassicAuthHandler(ctx)

	default:
		log.Printf("Endpoint (%s) not found.", pth)
	}
}
