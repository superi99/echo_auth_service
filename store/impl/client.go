package impl

import (
	"errors"
	"gorm.io/gorm"
	"skillspar/user_service/model"
)

type ClientStore struct {
	db *gorm.DB
}

func NewClientStore(db *gorm.DB) *ClientStore {
	return &ClientStore{
		db: db,
	}
}

func (cs *ClientStore) Create(client *model.Client) error {
	return cs.db.Create(client).Error
}

func (cs *ClientStore) FindByClientID(clientId string) (*model.Client, error) {
	client := new(model.Client)

	result := cs.db.First(&client, "id = ?", clientId)
	if errors.Is(gorm.ErrRecordNotFound, result.Error) {
		return nil, nil
	}
	return client, nil
}
