package jwt

import (
	"time"

	"github.com/valyala/fasthttp"
)

func DeleteRefreshTokenCookie(ctx *fasthttp.RequestCtx) {
	cookie := fasthttp.Cookie{}
	cookie.SetKey("refreshToken")
	cookie.SetPath("/")
	cookie.SetHTTPOnly(true)
	cookie.SetSecure(true)
	cookie.SetExpire(time.Now().Add(-1 * time.Hour)) // Set the time as a past time

	ctx.Response.Header.SetCookie(&cookie)
}
