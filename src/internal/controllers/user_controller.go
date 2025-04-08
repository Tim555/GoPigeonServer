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

func generateTag() string {

	return "aaaaa"
}

func CreateUser(user *models.User, db *gorm.DB) error {
	user.Tag = generateTag()
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func CreateContact(creator *models.User, concact *models.User, db *gorm.DB) error {

	contact := &models.Contact{
		CreatorID: creator.ID,
		ContactID: concact.ID,
	}
	err := db.Create(contact).Error
	if err != nil {
		return err
	}
	return nil
}

func GetContacts(user *models.User, db *gorm.DB) []models.Contact {
	contacts := []models.Contact{}
	db.Where("creator_id = ?", user.ID).Find(&contacts)

	return contacts
}
