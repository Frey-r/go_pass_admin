package controllers

import (
	"passcript/internal/utils"
	"path/filepath"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database_path string = "./internal/db/"

func Check_db() bool {
	db_path := "../db/"
	isDb, err := filepath.Match("*.db", db_path)
	if err != nil {
		utils.Log().Info("Error checking for database file", zap.Error(err))
		return isDb
	}
	return isDb
}

func Get_db() *gorm.DB {
	if !Check_db() {
		utils.Log().Info("Database file not found, creating...")
	}
	var db, err = gorm.Open(sqlite.Open(database_path+"passadmin.db"), &gorm.Config{})
	utils.Log().Info("Database succesfully created...")
	if err != nil {
		utils.Log().Error("Error opening database", zap.Error(err))
		return nil
	}
	return db
}
