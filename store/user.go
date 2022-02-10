package store

import "skillspar/user_service/model"

type UserStore interface {
	Create(user *model.User) error
	GetAll(offset, limit int) ([]model.User, int64, error)
	GetByEmail(email string) (*model.User, error)
	GetById(id uint) (*model.User, error)
}
