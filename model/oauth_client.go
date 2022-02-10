package model

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"time"
)

type Client struct {
	ID                   uuid.UUID `gorm:"type:char(36);primarykey"`
	UserId               int       `gorm:"index:idx_userId;"`
	Name                 string    `gorm:"not null"`
	Secret               string    `gorm:"size:100;"`
	Provider             string
	Redirect             string `gorm:"not null"`
	PersonalAccessClient bool   `gorm:"not null"`
	Revoked              bool   `gorm:"not null"`
}

func (client Client) BeforeCreate(tx *gorm.DB) error {
	tx.Statement.SetColumn("ID", uuid.New())
	return nil
}

type PersonalAccessClient struct {
	gorm.Model
	ClientId uuid.UUID `gorm:"type:char(36);"`
	Client   Client    `gorm:"references:ID"`
}

type AccessToken struct {
	gorm.Model
	UserId     int       `gorm:"index:idx_userId;"`
	ClientId   uuid.UUID `gorm:"type:char(36);"`
	Client     Client    `gorm:"references:ID"`
	Name       string
	Scopes     string
	Revoked    bool `gorm:"not null"`
	DeviceName string
	DeviceType string
	IpAddress  string `gorm:"type:char(15)"`
	expiresAt  time.Time
}

type RefreshToken struct {
	gorm.Model
	AccessTokenID uint        `gorm:"index:idx_accessTokenId;not null"`
	AccessToken   AccessToken `gorm:"references:ID"`
	Revoked       bool        `gorm:"not null"`
	ExpiresAt     time.Time
}
