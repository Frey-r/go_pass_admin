package controllers

import (
	"passcript/internal/models"
	"passcript/internal/utils"

	"go.uber.org/zap"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var database_path string = "./internal/db/"

func Get_db() *gorm.DB {
	var db, err = gorm.Open(sqlite.Open(database_path+"passadmin.db"), &gorm.Config{})
	if err != nil {
		utils.Log().Error("Error opening database", zap.Error(err))
		return nil
	}
	return db
}

func Migrate_tables() {
	db := Get_db()
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Pass{})
	db.Commit()
}
