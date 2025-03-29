package controllers

import (
	"passcript/internal/models"
	"passcript/internal/utils"

	"go.uber.org/zap"
)

func CreateUser(name string, password string) error {
	utils.Log().Info("Creating user")
	user := models.NewUser(name, password)
	db := GetDb()
	err := db.Create(user).Error
	if err != nil {
		utils.Log().Error("Error creating user", zap.Error(err))
		return err
	}
	return nil
}
