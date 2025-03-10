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

	// book handlers (CRUD) => table : book
	router.POST(dbApiMainPath+"/books/get/all", handlers.GetAllBooksHandler)          // get all books
	router.POST(dbApiMainPath+"/books/get/id/:id", handlers.GetBookByIDHandler)       // get book by id
	router.POST(dbApiMainPath+"/books/get/isbn/:isbn", handlers.GetBookByISBNHandler) // get book by ISBN
	router.POST(dbApiMainPath+"/books/create", handlers.CreateBookHandler)            // create book
	router.POST(dbApiMainPath+"/books/update", handlers.UpdateBookHandler)            // update book
	router.POST(dbApiMainPath+"/books/delete/id/:id", handlers.DeleteBookHandler)     // delete book by id

	// shelf_book handlers (CRUD) => table : shelf_book
	router.POST(dbApiMainPath+"/shelf-books/get/all", handlers.GetAllShelfBooksHandler)
	router.POST(dbApiMainPath+"/shelf-books/get/id/:id", handlers.GetShelfBookByIDHandler)
	router.POST(dbApiMainPath+"/shelf-books/create", handlers.CreateShelfBookHandler)
	router.POST(dbApiMainPath+"/shelf-books/delete/id/:id", handlers.DeleteShelfBookHandler)

	// user_shelf handlers (CRUD) => table : user_shelf
	router.POST(dbApiMainPath+"/user-shelves/get/all", handlers.GetAllUserShelvesHandler)
	router.POST(dbApiMainPath+"/user-shelves/get/id/:id", handlers.GetUserShelfByIDHandler)
	router.POST(dbApiMainPath+"/user-shelves/create", handlers.CreateUserShelfHandler)
	router.POST(dbApiMainPath+"/user-shelves/update", handlers.UpdateUserShelfHandler)
	router.POST(dbApiMainPath+"/user-shelves/delete/id/:id", handlers.DeleteUserShelfHandler)

	// follow handlers (CRUD) => table : follow
	router.POST(dbApiMainPath+"/follows/get/all", handlers.GetAllFollowsHandler)
	router.POST(dbApiMainPath+"/follows/get/id/:id", handlers.GetFollowByIDHandler)
	router.POST(dbApiMainPath+"/follows/create", handlers.CreateFollowHandler)
	router.POST(dbApiMainPath+"/follows/update", handlers.UpdateFollowHandler)
	router.POST(dbApiMainPath+"/follows/delete/id/:id", handlers.DeleteFollowHandler)

	// comment handlers (CRUD) => table : comment
	router.POST(dbApiMainPath+"/comments/get/all", handlers.GetAllCommentsHandler)
	router.POST(dbApiMainPath+"/comments/get/id/:id", handlers.GetCommentByIDHandler)
	router.POST(dbApiMainPath+"/comments/create", handlers.CreateCommentHandler)
	router.POST(dbApiMainPath+"/comments/update", handlers.UpdateCommentHandler)
	router.POST(dbApiMainPath+"/comments/delete/id/:id", handlers.DeleteCommentHandler)

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
