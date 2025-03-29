package controllers

import (
	"errors"
	"passcript/internal/models"
	"passcript/internal/utils"

	"go.uber.org/zap"
)

func isUserExist(name string) bool {
	db := GetDb()
	var user models.User
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		utils.Log().Error("Error checking user existence", zap.Error(err))
		return false
	}
	return true
}

func CreateUser(name string, password string) error {
	utils.Log().Info("Creating user" + name)
	if isUserExist(name) {
		utils.Log().Error("User already exists")
		return errors.New("user already exists")
	}
	encryptedPassword := Encrypter(GetPublicKey(), password)
	user := models.NewUser(name, encryptedPassword)
	db := GetDb()
	err := db.Create(user).Error
	if err != nil {
		utils.Log().Error("Error creating user", zap.Error(err))
		return err
	}
	db.Commit()
	utils.Log().Info("User  created successfully")
	return nil
}

func LoginUser(name string, password string) (error, int) {
	utils.Log().Info("Logging in user " + name)
	user := models.User{}
	db := GetDb()
	err := db.Where("name = ?", name).First(&user).Error
	if err != nil {
		utils.Log().Error("Error logging in user", zap.Error(err))
		return err, 0
	}
	encryptedPassword := user.GetPassword()
	decryptedPassword := Decrypter(GetPrivateKey(), encryptedPassword)
	if password != string(decryptedPassword) {
		utils.Log().Error("Error logging in user", zap.Error(err))
		return err, 0
	}
	utils.Log().Info("User logged in successfully")
	return nil, user.GetID()
}
