package main

import (
	"EC-site/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func router() http.Handler {
	// patternListenMux を返す
	mux := pat.New()
	// 各パスの設定・Get/Post 設定
	mux.Get("/login", http.HandlerFunc(handlers.LoginHandler))
	mux.Post("/create", http.HandlerFunc(handlers.CreateHandler))
	mux.Get("/delete", http.HandlerFunc(handlers.DeleteHandler))
	mux.Post("/deletePost", http.HandlerFunc(handlers.DeleteAccountHandler))
	mux.Get("/newResistration", http.HandlerFunc(handlers.NewResistrationHandler))
	mux.Post("/newResistrationPost", http.HandlerFunc(handlers.NewResistrationPostHandler))
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
