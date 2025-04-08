package controllers

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"gopigeon/internal/models"
)

var database *gorm.DB = nil
var migrated = false

func GetDB() (*gorm.DB, error) {
	if database == nil {
		var err error
		database, err = ConnectDB()
		if err != nil {
			return nil, err
		}
	}

	if !migrated {
		err := MigrateDB(database)
		if err != nil {
			return nil, err
		}
		migrated = true
	}

	return database, nil
}

func ConnectDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("chat.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func MigrateDB(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.Contact{})
	if err != nil {
		return err
	}
	return nil
}
