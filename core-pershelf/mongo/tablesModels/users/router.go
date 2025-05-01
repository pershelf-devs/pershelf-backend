package users

import (
	"log"
	"strings"

	"github.com/valyala/fasthttp"
)

type UserRouter interface {
	RegisterRoutes(router *fasthttp.RequestHandler)
}

type UserRouterImpl struct {
	controller UserController
}

func NewUserRouter(controller UserController) *UserRouterImpl {
	return &UserRouterImpl{controller: controller}
}

func (r *UserRouterImpl) RegisterRoutes(router *fasthttp.RequestHandler) {
	log.Printf("DEBUG: Registering user routes")

	handler := func(ctx *fasthttp.RequestCtx) {
		path := string(ctx.Path())
		method := string(ctx.Method())

		// ID parametresini çıkar
		if strings.HasPrefix(path, "/api/users/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 4 {
				ctx.SetUserValue("id", parts[3])
			}
		}

		// Username parametresini çıkar
		if strings.HasPrefix(path, "/api/users/username/") {
			parts := strings.Split(path, "/")
			if len(parts) >= 5 {
				ctx.SetUserValue("username", parts[4])
			}
		}

		switch {
		case path == "/api/users" && method == "POST":
			r.controller.CreateUser(ctx)
		case strings.HasPrefix(path, "/api/users/") && method == "GET":
			if strings.Contains(path, "/username/") {
				r.controller.GetUserByUsername(ctx)
			} else {
				r.controller.GetUserByID(ctx)
			}
		case strings.HasPrefix(path, "/api/users/") && method == "PUT":
			r.controller.UpdateUser(ctx)
		case strings.HasPrefix(path, "/api/users/") && method == "DELETE":
			r.controller.DeleteUser(ctx)
		default:
			ctx.SetStatusCode(fasthttp.StatusNotFound)
			ctx.SetBodyString(`{"error": "Not found"}`)
		}
	}

	*router = handler
}
