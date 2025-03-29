package models

import (
	"time"

	"gorm.io/gorm"
)

type Pass struct {
	gorm.Model
	id            string `gorm:"type:int;not null;primary_key;auto_increment"`
	serviceName   string `gorm:"type:varchar(100);not null"`
	userID        string `gorm:"type:varchar(100);not null;foreignKey:User(id)"`
	encryptedPass string `gorm:"type:varchar(150);not null"`
	createdAt     time.Time
	updatedAt     time.Time
}

func NewPass(serviceName string, userID string, encryptedPass string) *Pass {
	return &Pass{
		serviceName:   serviceName,
		userID:        userID,
		encryptedPass: encryptedPass,
	}
}

func (p *Pass) GetServiceName() string {
	return p.serviceName
}

func (p *Pass) SetServiceName(name string) {
	p.serviceName = name
}

func (p *Pass) GetEncryptedPass() string {
	return p.encryptedPass
}

func (p *Pass) SetEncryptedPass(encrypted string) {
	p.encryptedPass = encrypted
}

func (p *Pass) GetUserID() string {
	return p.userID
}

func (p *Pass) SetUserID(id string) {
	p.userID = id
}

func (p *Pass) GetCreatedAt() time.Time {
	return p.createdAt
}

func (p *Pass) SetCreatedAt(createdAt time.Time) {
	p.createdAt = createdAt
}

func (p *Pass) GetUpdatedAt() time.Time {
	return p.updatedAt
}

func (p *Pass) SetUpdatedAt(updatedAt time.Time) {
	p.updatedAt = updatedAt
}

func (p *Pass) GetID() string {
	return p.id
}
