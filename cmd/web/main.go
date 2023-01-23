package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := router()
	log.Println("Starting channel listener...")
	fmt.Println("Server starting up ...... localhost:8080/login")
	_ = http.ListenAndServe(":8080", mux)
}
