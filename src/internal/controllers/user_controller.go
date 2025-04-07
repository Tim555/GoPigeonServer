package controllers

import (
	"gopigeon/internal/models"
)

func GetUser(username string) (*models.User, error) {

	var testUser = models.User{
		Username: username,
		Password: "testpassword",
		Tag:      "testtag",
	}
	return &testUser, nil
}
