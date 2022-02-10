package handler

import (
	"github.com/labstack/echo/v4"
	EchoMiddleware "github.com/labstack/echo/v4/middleware"
	"skillspar/user_service/server/middleware"
	"skillspar/user_service/utils"
)

func (h *Handler) Register(v1 *echo.Group) {

	// middleware

	// routes
	api := v1.Group("/api")

	guestUsers := api.Group("/users")
	guestUsers.POST("/signup", h.Signup)
	guestUsers.POST("/login", h.Login)

	user := api.Group("/user")
	user.GET("/whoami", h.Whoami, EchoMiddleware.JWTWithConfig(EchoMiddleware.JWTConfig{
		SigningKey:  utils.JWTSecret,
		TokenLookup: "header:Authorization", // as default
		AuthScheme:  "Bearer",               // as default
	}))

	//user := v1.Group("/user") // add middleware
	//user.GET("", h.CurrentUser)
	//user.PUT("", h.UpdateUser)

	oauth2 := v1.Group("/oauth")

	authorize := oauth2.Group("/authorize", middleware.GuestMiddleware)
	authorize.GET("", h.Authorize, middleware.CSRFMiddleware())
	authorize.POST("", h.Approve)
	//authorize.DELETE("", h.Authorize)
	authorize.GET("/token", h.IssueToken)

	oauth2.GET("/create-client", h.CreateClient)
	oauth2.GET("/login", h.GetLogin, middleware.CSRFMiddleware())
	oauth2.POST("/login", h.PostLogin)

}
