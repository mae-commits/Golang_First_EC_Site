package main

import (
	"EC-site/internal/handlers"
	"net/http"

	"github.com/bmizerany/pat"
)

func router() http.Handler {
	mux := pat.New()
	mux.Get("/login", http.HandlerFunc(handlers.LoginHandler))
	mux.Post("/create", http.HandlerFunc(handlers.CreateHandler))
	mux.Post("/delete", http.HandlerFunc(handlers.DeleteHandler))
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Get("/static/", http.StripPrefix("/static", fileServer))
	return mux
}
