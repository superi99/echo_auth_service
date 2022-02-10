package handler

import (
	"github.com/labstack/echo/v4"
	"skillspar/user_service/model"
)

type SigninRequest struct {
	Email      string `json:"email" validate:"required,email"`
	Password   string `json:"password" validate:"required,password"`
	RememberMe bool   `json:"rememberMe"`
}

func (r *SigninRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}

type SignupRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,password"`
	Username        string `json:"username"`
	ConfirmPassword string `json:"confirmPassword" validate:"eqfield=Password"`
	//FirstName string `json:"firstName"`
	//LastName  string `json:"LastName" `
	//NickName  string `json:"nickName"`
	//ClientId  string `json:"clientId" validate:"required"`
}

func (r *SignupRequest) bind(c echo.Context, u *model.User) error {
	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}
	u.Email = r.Email
	hPassword, err := u.HashPassword(r.Password)
	if err != nil {
		return err
	}
	u.Password = hPassword
	u.Username = r.Username
	return nil
}

/** oauth2 **/
type AuthorizeRequest struct {
	ClientID     string `query:"clientId" validate:"required,uuid"`
	RedirectUri  string `query:"redirectUri" validate:"required"`
	ResponseType string `query:"responseType" validate:"required"`
}

func (r *AuthorizeRequest) Bind(c echo.Context) error {

	if err := c.Bind(r); err != nil {
		return err
	}
	if err := c.Validate(r); err != nil {
		return err
	}

	return nil
}

type LoginRequest struct {
	Email      string `form:"email" validate:"required,email"`
	Password   string `form:"password" validate:"required,password"`
	RememberMe bool   `form:"rememberMe"`
}

func (r *LoginRequest) bind(c echo.Context) error {
	if err := c.Bind(r); err != nil {
		return err
	}

	if err := c.Validate(r); err != nil {
		return err
	}
	return nil
}
