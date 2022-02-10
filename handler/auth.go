package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"net/url"
	"skillspar/user_service/utils"
	"strings"
)

func (h *Handler) Authorize(c echo.Context) error {

	// example for all scopes

	authToken := utils.RandStringRunes(40)

	sess, _ := utils.GetSession(c)
	sess.Values["authToken"] = authToken
	sess.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "authorize.html", map[string]interface{}{
		"csrf":      c.Get("csrf").(string),
		"authToken": authToken,
	})
}

func (h *Handler) Approve(c echo.Context) error {
	sess, _ := utils.GetSession(c)
	if authToken := c.FormValue("authToken"); authToken != "" && authToken == sess.Values["authToken"] {

		code := utils.RandStringRunes(40)
		sess.Values["authCode"] = code
		sess.Save(c.Request(), c.Response())

		return c.JSON(http.StatusOK, map[string]interface{}{
			"code": code,
		})

	}
	return c.JSON(http.StatusUnauthorized, "Wrong auth token")
}

func (h *Handler) IssueToken(c echo.Context) error {
	sess, _ := utils.GetSession(c)

	if authCode := c.QueryParam("code"); authCode != "" && authCode == sess.Values["authCode"] {
		userEmail := sess.Values["user"].(string)
		user, _ := h.UserStore.GetByEmail(userEmail)
		return c.JSON(http.StatusOK, newUserResponse(user))
	}
	return c.JSON(http.StatusUnauthorized, "Wrong authorization code")
}

func (h *Handler) GetLogin(c echo.Context) error {

	req := &AuthorizeRequest{}
	if err := req.Bind(c); err != nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NewError(err))
	}

	// check client credential
	client, err := h.ClientStore.FindByClientID(req.ClientID)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.NewError(err))
	}
	if client == nil {
		return c.JSON(http.StatusUnprocessableEntity, utils.NotFound())
	}

	redirectUrls := strings.Split(client.Redirect, ",")
	if !utils.Contains(redirectUrls, req.RedirectUri) {
		return c.JSON(http.StatusUnprocessableEntity, utils.NotFound())
	}

	// check client credential successfully
	// add client credential to session
	sess, _ := utils.GetSession(c)
	sess.Values["clientId"] = req.ClientID
	sess.Values["redirectUri"] = req.RedirectUri
	sess.Values["responseType"] = req.ResponseType
	sess.Save(c.Request(), c.Response())

	return c.Render(http.StatusOK, "login.html", map[string]interface{}{
		"csrf": c.Get("csrf").(string),
	})
}

func (h *Handler) PostLogin(c echo.Context) error {

	req := &LoginRequest{}
	if err := req.bind(c); err != nil {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"csrf": c.Get("csrf").(string),
			"msg":  "Failed to validation",
		})
	}

	u, err := h.UserStore.GetByEmail(req.Email)
	if err != nil {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"csrf": c.Get("csrf").(string),
			"msg":  "Something wrong",
		})
	}
	if u == nil {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"csrf": c.Get("csrf").(string),
			"msg":  "Email does not exists",
		})
	}
	if !u.CheckPassword(req.Password) {
		return c.Render(http.StatusOK, "login.html", map[string]interface{}{
			"csrf": c.Get("csrf").(string),
			"msg":  "Username or password are incorrect",
		})

	}

	sess, _ := utils.GetSession(c)
	// store use to the session
	sess.Values["user"] = u.Email
	// redirect to authorize
	sess.Save(c.Request(), c.Response())

	query := url.Values{
		"clientId":     {sess.Values["clientId"].(string)},
		"redirectUri":  {sess.Values["redirectUri"].(string)},
		"responseType": {sess.Values["responseType"].(string)},
		"scope":        {""},
	}
	return c.Redirect(http.StatusSeeOther, "/oauth/authorize?"+query.Encode())

}
