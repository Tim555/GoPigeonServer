package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID       uint   `json:"id" gorm:"primaryKey"`
	Username string `json:"username" gorm:"unique;not null"`
	Password string `json:"password" gorm:"not null"`

	Tag string `json:"tag" gorm:"not null"`
}

type Contact struct {
	gorm.Model
	ID uint `json:"id" gorm:"primaryKey"`

	CreatorID uint `json:"creator_id" gorm:"not null"`
	Creator   User `json:"creator" gorm:"foreignKey:CreatorID"`

	ContactID uint `json:"contact_id" gorm:"not null"`
	Contact   User `json:"contact" gorm:"foreignKey:ContactID"`
}
