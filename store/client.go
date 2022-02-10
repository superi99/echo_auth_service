package store

import "skillspar/user_service/model"

type ClientStore interface {
	Create(client *model.Client) error
	FindByClientID(clientId string) (*model.Client, error)
}
