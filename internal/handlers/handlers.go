package handlers

import (
	"EC-site/domain"
	"EC-site/internal/accounts"
	"log"
	"net/http"
	"strings"
	"text/template"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// ログインページ
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "html/loginPage.html")
	if err != nil {
		log.Fatal(err)
	}
}

// ログイン認証処理
func CreateHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("userName")
	password := r.FormValue("password")
	isUserName := strings.ReplaceAll(userName, " ", "")
	isPassword := strings.ReplaceAll(password, " ", "")
	if isUserName == "" || isPassword == "" {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
	count := accounts.Authentification(userName, password)
	if count != 0 {
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// ユーザ情報を削除する処理
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("userName")
	password := r.FormValue("password")
	isUserName := strings.ReplaceAll(userName, " ", "")
	isPassword := strings.ReplaceAll(password, " ", "")
	db, err := gorm.Open(sqlite.Open("userData.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.User{})
	if isUserName == "" || isPassword == "" {
		http.Redirect(w, r, "/delete", http.StatusFound)
	}
	count := accounts.Authentification(userName, password)
	if count != 0 {
		db.Where("username = ?", userName).Where("password = ?", password).Delete(&domain.User{})
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Redirect(w, r, "/delete", http.StatusFound)
	}
}

// ページを開く処理
func renderPage(w http.ResponseWriter, html string) error {
	_, err := template.ParseFiles(html)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}
