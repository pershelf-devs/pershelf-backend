package starter

import (
	"log"

	"github.com/core-pershelf/globals"
	"github.com/valyala/fasthttp"
)

func StartServer(srv *fasthttp.Server) {
	// Log the server's starting message
	log.Printf("Server started listening on %s:%s", globals.ServerConf.Server.ServerIP, globals.ServerConf.Server.ServerPort)

	// Start the server
	if err := srv.ListenAndServe(globals.ServerConf.Server.ServerIP + ":" + globals.ServerConf.Server.ServerPort); err != nil {
		log.Fatalf("(Error): error running the server : %v", err)
	}

	// log the server's graceful closing message
	log.Println("Server is closed gracefully.")
}
