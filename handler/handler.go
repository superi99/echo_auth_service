package handler

import (
	"skillspar/user_service/store"
)

type Handler struct {
	UserStore   store.UserStore
	ClientStore store.ClientStore
}

func NewHandler(us store.UserStore, cs store.ClientStore) *Handler {
	return &Handler{
		UserStore:   us,
		ClientStore: cs,
	}
}
