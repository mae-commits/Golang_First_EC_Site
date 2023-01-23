package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	userName string
	password string
}
