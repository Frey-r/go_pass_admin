package models

import (
	"time"

	"gorm.io/gorm"
)

type pass struct {
	gorm.Model
	id            string `gorm:"type:int;not null;primary_key;auto_increment"`
	serviceName   string `gorm:"type:varchar(100);not null"`
	userID        string `gorm:"type:varchar(100);not null;foreignKey:User(id)"`
	encryptedPass string `gorm:"type:varchar(150);not null"`
	createdAt     time.Time
	updatedAt     time.Time
}

func NewPass(serviceName string, userID string, encryptedPass string) *pass {
	return &pass{
		serviceName:   serviceName,
		userID:        userID,
		encryptedPass: encryptedPass,
	}
}

func (p *pass) GetServiceName() string {
	return p.serviceName
}

func (p *pass) SetServiceName(name string) {
	p.serviceName = name
}

func (p *pass) GetEncryptedPass() string {
	return p.encryptedPass
}

func (p *pass) SetEncryptedPass(encrypted string) {
	p.encryptedPass = encrypted
}

func (p *pass) GetUserID() string {
	return p.userID
}

func (p *pass) SetUserID(id string) {
	p.userID = id
}

func (p *pass) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *pass) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p *pass) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p *pass) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

func (p *pass) GetID() string {
	return p.id
}
