package accounts

import (
	"EC-site/domain"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ユーザ認証
func Authentification(userName string, password string) (count int64) {
	db, err := gorm.Open(sqlite.Open("userData.db"), &gorm.Config{})
	db.AutoMigrate(&domain.User{})
	if err != nil {
		panic("failed to connect database.")
	}
	db.AutoMigrate(&domain.User{})
	db.Model(&domain.User{}).Where("username = ?", userName).Where("password = ?", password).Count(&count)
	return count
}
