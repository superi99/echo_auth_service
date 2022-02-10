package server

import (
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
	"skillspar/user_service/utils"
)

func New() *echo.Echo {
	e := echo.New()
	e.Logger.SetLevel(log.DEBUG)
	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `"error":"${error}", "uri""="${uri}"\n`,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept, echo.HeaderAuthorization},
		AllowMethods: []string{echo.GET, echo.HEAD, echo.PUT, echo.PATCH, echo.POST, echo.DELETE},
	}))
	//e.Use(middleware.CSRFWithConfig(middleware.CSRFConfig{
	//	TokenLength:  32,
	//	TokenLookup:  "form:csrf",
	//	ContextKey:   "csrf",
	//	CookieName:   "_csrf",
	//	CookieMaxAge: 86400,
	//}))

	// session config
	var sessionStore = sessions.NewCookieStore([]byte(utils.Getenv("SESSION_SECRET", "SECRET")))
	sessionStore.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	e.Use(session.Middleware(sessionStore))

	e.Validator = NewValidator()
	return e
}
