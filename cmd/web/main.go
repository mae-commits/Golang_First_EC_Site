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
	//  サーバを:8080ポートで起動
	_ = http.ListenAndServe(":8080", mux)
}
