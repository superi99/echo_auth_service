package handler

import (
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"net/http"
	"skillspar/user_service/model"
	"skillspar/user_service/utils"
)

func (h *Handler) Signup(c echo.Context) error {
	var u = new(model.User)
	req := &SignupRequest{}
	if err := req.bind(c, u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	if err := h.UserStore.Create(u); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	return c.JSON(http.StatusOK, newUserResponse(u))

}

func (h *Handler) Login(c echo.Context) error {
	req := &SigninRequest{}
	if err := req.bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	u, err := h.UserStore.GetByEmail(req.Email)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if u == nil {
		return c.JSON(http.StatusUnauthorized, utils.EmailPasswordUnAuthorized())
	}
	if !u.CheckPassword(req.Password) {
		return c.JSON(http.StatusUnauthorized, utils.EmailPasswordUnAuthorized())
	}

	return c.JSON(http.StatusOK, newUserResponse(u))
}

func (h *Handler) CreateClient(c echo.Context) error {
	//clientStore := store
	return nil
}

func (h *Handler) Whoami(c echo.Context) error {
	token := c.Get("user").(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	userId := uint(claims["id"].(float64))
	user, err := h.UserStore.GetById(userId)
	if err != nil {
		return c.JSON(http.StatusNotFound, "not found user")
	}

	return c.JSON(http.StatusOK, newUserResponse(user))
}
