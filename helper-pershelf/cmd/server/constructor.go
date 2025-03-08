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

	// user_book handlers (CRUD) => table : user_book
	router.POST(dbApiMainPath+"/user-books/get/all", handlers.GetAllUserBooksHandler)                   // get all user books
	router.POST(dbApiMainPath+"/user-books/get/id/:id", handlers.GetUserBookByIDHandler)                // get user book by id
	router.POST(dbApiMainPath+"/user-books/get/user-id/:user-id", handlers.GetUserBooksByUserIDHandler) // get user books by user id
	router.POST(dbApiMainPath+"/user-books/get/book-id/:book-id", handlers.GetUserBooksByBookIDHandler) // get user books by book id
	router.POST(dbApiMainPath+"/user-books/create", handlers.CreateUserBookHandler)                     // create user book entry
	router.POST(dbApiMainPath+"/user-books/update", handlers.UpdateUserBookHandler)                     // update user book entry
	router.POST(dbApiMainPath+"/user-books/delete/id/:id", handlers.DeleteUserBookHandler)              // delete user book entry by id

	// review handlers (CRUD) => table : review
	router.POST(dbApiMainPath+"/reviews/get/all", handlers.GetAllReviewsHandler)                   // get all reviews
	router.POST(dbApiMainPath+"/reviews/get/id/:id", handlers.GetReviewByIDHandler)                // get review by id
	router.POST(dbApiMainPath+"/reviews/get/user-id/:user-id", handlers.GetReviewsByUserIDHandler) // get reviews by user id
	router.POST(dbApiMainPath+"/reviews/get/book-id/:book-id", handlers.GetReviewsByBookIDHandler) // get reviews by book id
	router.POST(dbApiMainPath+"/reviews/create", handlers.CreateReviewHandler)                     // create review entry
	router.POST(dbApiMainPath+"/reviews/update", handlers.UpdateReviewHandler)                     // update review entry
	router.POST(dbApiMainPath+"/reviews/delete/id/:id", handlers.DeleteReviewHandler)              // delete review entry by id

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
