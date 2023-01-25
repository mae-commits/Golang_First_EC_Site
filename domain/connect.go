package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	UserName string
	Password string
}
