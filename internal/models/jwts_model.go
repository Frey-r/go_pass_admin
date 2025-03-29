package models

import (
	"time"

	"gorm.io/gorm"
)

type JWT struct {
	gorm.Model
	id        int    `gorm:"type:int;not null;primary_key;auto_increment"`
	UserID    int    `gorm:"type:int;not null"`
	Token     string `gorm:"type:text;not null"`
	ExpiresAt time.Time
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewJWT(userID int, token string) *JWT {
	return &JWT{
		UserID:    userID,
		Token:     token,
		ExpiresAt: time.Now().Add(time.Hour * 24 * 7),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Model:     gorm.Model{},
	}
}

func (j *JWT) GetUserID() int {
	return j.UserID
}

func (j *JWT) SetUserID(userID int) {
	j.UserID = userID
}

func (j *JWT) GetToken() string {
	return j.Token
}

func (j *JWT) SetToken(token string) {
	j.Token = token
}

func (j *JWT) GetExpiresAt() time.Time {
	return j.ExpiresAt
}

func (j *JWT) SetExpiresAt(years int) {
	expiresAt := time.Now().Add(time.Hour * 24 * 365 * time.Duration(years))
	j.ExpiresAt = expiresAt
}

func (j *JWT) GetCreatedAt() time.Time {
	return j.CreatedAt
}

func (j *JWT) SetCreatedAt(createdAt time.Time) {
	j.CreatedAt = createdAt
}

func (j *JWT) GetUpdatedAt() time.Time {
	return j.UpdatedAt
}

func (j *JWT) SetUpdatedAt(updatedAt time.Time) {
	j.UpdatedAt = updatedAt
}

func (j *JWT) GetID() int {
	return j.id
}
