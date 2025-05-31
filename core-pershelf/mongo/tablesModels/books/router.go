package books

import (
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

type BookRouter interface {
	RegisterRoutes(router *fasthttp.RequestHandler)
}

type BookRouterImpl struct {
	controller BookController
}

func NewBookRouter(controller BookController) *BookRouterImpl {
	return &BookRouterImpl{controller: controller}
}

func (r *BookRouterImpl) RegisterRoutes(router *fasthttp.RequestHandler) {
	log.Printf("DEBUG: Registering book routes")

	handler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		method := string(ctx.Method())

		// ID parametresini çıkar
		if strings.HasPrefix(path, "/api/books/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 4 {
				ctx.SetUserValue("id", parts[3])
			}
		}

		// Owner ID parametresini çıkar
		if strings.HasPrefix(path, "/api/books/owner/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 5 {
				ctx.SetUserValue("owner_id", parts[4])
			}
		}

		// Status parametresini çıkar
		if strings.HasPrefix(path, "/api/books/status/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 5 {
				ctx.SetUserValue("status", parts[4])
			}
		}

		switch {
		case path == "/api/books" && method == "POST":
			r.controller.CreateBook(ctx)
		case strings.HasPrefix(path, "/api/books/") && method == "GET":
			if strings.Contains(path, "/owner/") {
				r.controller.GetBooksByOwnerID(ctx)
			} else if strings.Contains(path, "/status/") {
				r.controller.GetBooksByStatus(ctx)
			} else {
				r.controller.GetBookByID(ctx)
			}
		case strings.HasPrefix(path, "/api/books/") && method == "PUT":
			r.controller.UpdateBook(ctx)
		case strings.HasPrefix(path, "/api/books/") && method == "DELETE":
			r.controller.DeleteBook(ctx)
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(`{"error": "Not found"}`)
		}
	}

	*router = handler
}
