package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	ID        uint
	FirstName string
	LastName  string
	Email     string `gorm:"unique"`
	Password  string
}
