package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	id        int    `gorm:"type:int;not null;primary_key;auto_increment"`
	Name      string `gorm:"type:varchar(100);not null"`
	Password  []byte `gorm:"type:blob;not null"`
	createdAt time.Time
	updatedAt time.Time
}

func NewUser(name string, password []byte) *User {
	return &User{
		Name:      name,
		Password:  password,
		createdAt: time.Now(),
		updatedAt: time.Now(),
		Model:     gorm.Model{},
	}
}
func (u *User) GetName() string {
	return u.Name
}

func (u *User) SetName(name string) {
	u.Name = name
}

func (u *User) GetPassword() []byte {
	return u.Password
}

func (u *User) SetPassword(password []byte) {
	u.Password = password
}

func (u *User) GetCreatedAt() time.Time {
	return u.createdAt
}

func (u *User) SetCreatedAt(createdAt time.Time) {
	u.createdAt = createdAt
}

func (u *User) GetUpdatedAt() time.Time {
	return u.updatedAt
}

func (u *User) SetUpdatedAt(updatedAt time.Time) {
	u.updatedAt = updatedAt
}

func (u *User) GetID() int {
	return u.id
}
