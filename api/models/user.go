package models

import "gorm.io/gorm"

type User struct {
	gorm.Model //created_atなど、その他のDBデータも追加してくれる
	ID         uint
	FirstName  string
	LastName   string
	Email      string `gorm:"unique"`
	Password   []byte
}
