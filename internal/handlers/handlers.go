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

type loginUser struct {
	userName string
	password string
}

// ログインページ
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "html/loginPage.html", "")
	if err != nil {
		log.Fatal(err)
	}
}

// ユーザ新規登録ページ
func NewResistrationHandler(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "html/newResistrationPage.html", "")
	if err != nil {
		log.Fatal(err)
	}
}

//ユーザ情報削除ページ
func DeleteHandler(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "html/deleteAccountPage.html", "")
	if err != nil {
		log.Fatal(err)
	}
}

// 商品紹介ページ
func MainHandler(w http.ResponseWriter, r *http.Request) {
	err := renderPage(w, "html/mainPage.html", nowLoginUser.userName)
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
		loginUserNow := New(userName, password)
		http.Redirect(w, r, "/main", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

//ログインユーザ名を取得する処理
func New(userName string, password string) *loginUser {
	return &loginUser{userName: userName, password: password}
}

// ユーザ情報を削除する処理
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
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
		db.Where("user_name = ?", userName).Where("password = ?", password).Delete(&domain.User{})
		http.Redirect(w, r, "/login", http.StatusFound)
	} else {
		http.Redirect(w, r, "/delete", http.StatusFound)
	}
}

// 新たなアカウントを登録する処理
func NewResistrationPostHandler(w http.ResponseWriter, r *http.Request) {
	userName := r.FormValue("userName")
	password := r.FormValue("password")
	isUserName := strings.ReplaceAll(userName, " ", "")
	isPassword := strings.ReplaceAll(password, " ", "")
	if isUserName == "" || isPassword == "" {
		http.Redirect(w, r, "/newResistration", http.StatusFound)
	}
	db, err := gorm.Open(sqlite.Open("userData.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	db.AutoMigrate(&domain.User{})
	count := accounts.Authentification(userName, password)
	if count != 0 {
		http.Redirect(w, r, "/newResistration", http.StatusFound)
	} else {
		db.Create(&domain.User{UserName: userName, Password: password})
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}

// ページを開く処理
func renderPage(w http.ResponseWriter, html string, userName string) error {
	page, err := template.ParseFiles(html)
	if err != nil {
		log.Fatal(err)
		return err
	}
	err = page.Execute(w, userName)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
