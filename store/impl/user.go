package impl

import (
	"gorm.io/gorm"
	"skillspar/user_service/model"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (us *UserStore) Create(user *model.User) error {
	return us.db.Create(user).Error
}

func (us *UserStore) GetAll(offset, limit int) ([]model.User, int64, error) {
	var (
		users []model.User
		count int64
	)

	//us.db.Model(&users).Count(&count)
	//us.db.Offset(offset).
	//	Limit(limit).
	//	Order("created_at desc").Find(&users)
	us.db.Find(&users)
	return users, count, nil

}

func (us *UserStore) GetByEmail(email string) (*model.User, error) {
	var m model.User

	err := us.db.Where(&model.User{
		Email: email,
	}).First(&m).Error

	if err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}

	return &m, nil
}

func (us *UserStore) GetById(id uint) (*model.User, error) {
	var m model.User
	if err := us.db.First(&m, id).Error; err != nil {
		if gorm.ErrRecordNotFound == err {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}
