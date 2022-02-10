package middleware

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"net/http"
	"skillspar/user_service/utils"
)

func GuestMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		sess, _ := utils.GetSession(c)
		user := sess.Values["user"]
		if user == nil {
			return c.Redirect(http.StatusPermanentRedirect, "/oauth/login?"+c.QueryString())
		}
		return next(c)
	}
}

func CSRFMiddleware() echo.MiddlewareFunc {
	return middleware.CSRFWithConfig(middleware.CSRFConfig{
		TokenLength:  32,
		TokenLookup:  "form:csrf",
		ContextKey:   "csrf",
		CookieName:   "_csrf",
		CookieMaxAge: 86400,
	})
}
