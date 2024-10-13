package database

import (
	"main/model"

	"gorm.io/gorm"
)

func GetUserFromEmail(db *gorm.DB) model.User {
	var user model.User
	db.Where("email = ?", "me@gmail.com").Find(&user)

	return user
}
