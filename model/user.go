package model

import (
	"gorm.io/gorm"
)

// User model
type User struct {
	gorm.Model
	Name      string
	IpAddress string
	Email     string
	Password  string
}
