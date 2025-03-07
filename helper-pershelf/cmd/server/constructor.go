package server

import (
	"fmt"
	"log"

	"github.com/buaazp/fasthttprouter"
	config2 "github.com/pershelf/pershelf/config/server"
	"github.com/pershelf/pershelf/endpoint/handlers"
	"github.com/pershelf/pershelf/globals"
	"github.com/valyala/fasthttp"
)

const (
	dbApiMainPath = "/restapi/helper/v1.0"
)

func RunDBHttpServer(conf config2.ServerConfig) error {
	router := fasthttprouter.New()

	router.POST(dbApiMainPath+"/test", func(ctx *fasthttp.RequestCtx) {
		globals.Log("Test endpoint hit successfully")
	})

	// users handlers (CRUD) => table : users
	router.POST(dbApiMainPath+"/users/get/all", handlers.GetAllUsersHandler)      // get all users
	router.POST(dbApiMainPath+"/users/get/id/:id", handlers.GetUserByIDHandler)   // get user by id
	router.POST(dbApiMainPath+"/users/create", handlers.CreateUserHandler)        // create user
	router.POST(dbApiMainPath+"/users/update", handlers.UpdateUserHandler)        // update user
	router.POST(dbApiMainPath+"/users/delete/id/:id", handlers.DeleteUserHandler) // delete user by id

	// refresh tokens handlers (CRUD) => table : refresh tokens
	router.POST(dbApiMainPath+"/refresh-tokens/get/all", handlers.GetAllRefreshTokensHandler)                  // get all refresh tokens
	router.POST(dbApiMainPath+"/refresh-tokens/get/id/:id", handlers.GetRefreshTokenByIDHandler)               // get refresh token by id
	router.POST(dbApiMainPath+"/refresh-tokens/get/user-id/:user-id", handlers.GetRefreshTokenByUserIDHandler) // get refresh token by id
	router.POST(dbApiMainPath+"/refresh-tokens/create", handlers.CreateRefreshTokenHandler)                    // create refresh token
	router.POST(dbApiMainPath+"/refresh-tokens/update", handlers.UpdateRefreshTokenHandler)                    // update refresh token
	router.POST(dbApiMainPath+"/refresh-tokens/delete/id/:id", handlers.DeleteRefreshTokenHandler)             // delete refresh token by id

	srv := &fasthttp.Server{
		Handler: router.Handler,
	}

	// set the read buffer size to 10000
	srv.ReadBufferSize = 10000

	// set the write buffer size to 10000
	srv.WriteBufferSize = 10000

	log.Printf("Server started listening on port :  %s", conf.Port)
	if err := srv.ListenAndServe(fmt.Sprintf(":%s", conf.Port)); err != nil {
		globals.Log("Error starting the server: ", err)
		return err
	}
	log.Printf("Server stopped listening on port :  %s", conf.Port)

	return nil
}
