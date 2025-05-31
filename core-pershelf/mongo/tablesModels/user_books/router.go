package user_books

import (
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

type UserBookRouter interface {
	RegisterRoutes(router *fasthttp.RequestHandler)
}

type UserBookRouterImpl struct {
	controller UserBookController
}

func NewUserBookRouter(controller UserBookController) *UserBookRouterImpl {
	return &UserBookRouterImpl{controller: controller}
}

func (r *UserBookRouterImpl) RegisterRoutes(router *fasthttp.RequestHandler) {
	log.Printf("DEBUG: Registering user_books routes")

	handler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		method := string(ctx.Method())

		// ID parametresini çıkar
		if strings.HasPrefix(path, "/api/user_books/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 4 {
				ctx.SetUserValue("id", parts[3])
			}
		}

		// UserID parametresini çıkar
		if strings.HasPrefix(path, "/api/user_books/user/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 5 {
				ctx.SetUserValue("user_id", parts[4])
			}
		}

		switch {
		case path == "/api/user_books" && method == "POST":
			r.controller.CreateUserBook(ctx)
		case strings.HasPrefix(path, "/api/user_books/") && method == "GET":
			if strings.Contains(path, "/user/") {
				r.controller.GetUserBooksByUserID(ctx)
			} else {
				r.controller.GetUserBookByID(ctx)
			}
		case strings.HasPrefix(path, "/api/user_books/") && method == "PUT":
			r.controller.UpdateUserBook(ctx)
		case strings.HasPrefix(path, "/api/user_books/") && method == "DELETE":
			r.controller.DeleteUserBook(ctx)
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(`{"error": "Not found"}`)
		}
	}

	*router = handler
}
