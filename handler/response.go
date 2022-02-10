package handler

import (
	"skillspar/user_service/model"
	"skillspar/user_service/utils"
)

type UserResponse struct {
	Email    string `json:"email"`
	Username string `json:"username"`
	Token    string `json:"token"`
}

func newUserResponse(u *model.User) *UserResponse {
	r := new(UserResponse)
	r.Username = u.Username
	r.Email = u.Email
	r.Token = utils.GenerateJWT(u.ID)
	return r
}
