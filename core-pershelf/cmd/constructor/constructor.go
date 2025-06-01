package constructor

import (
	"log"

	"github.com/core-pershelf/endpoints/handlers/auth"
	"github.com/core-pershelf/endpoints/handlers/bookLikes"
	"github.com/core-pershelf/endpoints/handlers/books"
	"github.com/core-pershelf/endpoints/handlers/reviews"
	"github.com/core-pershelf/endpoints/handlers/test" // <-- added this line
	"github.com/core-pershelf/endpoints/handlers/users"
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
	const (
		apiPathHeader = "/restapi/v1.0"
	)

	switch pth := string(ctx.Path()); pth {

	case "/test":
		test.ExecuteTestHandler(ctx)

	case apiPathHeader + "/auth/login":
		auth.ClassicAuthHandler(ctx)
	case apiPathHeader + "/auth/register":
		auth.UserRegisterHandler(ctx)

	case apiPathHeader + "/books/discover/most-reads":
		books.GetMostReadBooksHandler(ctx)
	case apiPathHeader + "/dashboard/user/recommended-books":
		books.GetUserRecomendedBooksHandler(ctx)

	case apiPathHeader + "/books/create":
		books.CreateBookHandler(ctx)

	case apiPathHeader + "/books/get/id":
		books.GetBookByIDHandler(ctx)

	case apiPathHeader + "/books/get/by-genre":
		books.GetBooksByGenreHandler(ctx)

	case apiPathHeader + "/users/get/id":
		users.GetUserByIDHandler(ctx)
	case apiPathHeader + "/users/update/profile-photo":
		users.UpdateUserProfilePhotoHandler(ctx)

	case apiPathHeader + "/book-likes/like/by-book-id":
		bookLikes.CreateBookLikeHandler(ctx)

	case apiPathHeader + "/reviews/get/book-reviews":
		reviews.GetReviewsByBookIDHandler(ctx)
	case apiPathHeader + "/reviews/create/book-review":
		reviews.CreateBookReviewHandler(ctx)

	case apiPathHeader + "/reviews/get/by-user":
		reviews.GetReviewsByUserIDHandler(ctx)

	default:
		log.Printf("Endpoint (%s) not found.", pth)
	}
}
