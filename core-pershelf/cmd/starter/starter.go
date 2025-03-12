package starter

import (
	"log"

	"github.com/valyala/fasthttp"
)

func StartServer(srv *fasthttp.Server) {
	// Log the server's starting message
	log.Printf("Server started listening on port : 443")

	// Start the server
	if err := srv.ListenAndServe("127.0.0.1:443"); err != nil {
		log.Fatalf("(Error): error running the server : %v", err)
	}

	// log the server's graceful closing message
	log.Println("Server is closed gracefully.")
}