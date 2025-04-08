package controllers

import (
	"gopigeon/internal/models"

	"gorm.io/gorm"
)

func GetUser(username string, db *gorm.DB) *models.User {

	user := &models.User{}
	db.Where("username = ?", username).First(user)

	return user
}

func GetUsers(db *gorm.DB) []models.User {
	users := []models.User{}
	db.Find(&users)

	return users
}

func CreateUser(user *models.User, db *gorm.DB) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}
