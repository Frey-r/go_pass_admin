package models

import (
	"time"

	"gorm.io/gorm"
)

type user struct {
	gorm.Model
	id        string `gorm:"type:int;not null;primary_key;auto_increment"`
	name      string `gorm:"type:varchar(100);not null"`
	password  string `gorm:"type:varchar(150);not null"`
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name string, password string) *user {
	return &user{
		name:     name,
		password: password,
	}
}
func (u *user) GetName() string {
	return u.name
}

func (u *user) SetName(name string) {
	u.name = name
}

func (u *user) GetPassword() string {
	return u.password
}

func (u *user) SetPassword(password string) {
	u.password = password
}

func (u *user) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *user) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *user) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *user) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

func (u *user) GetID() string {
	return u.id
}
