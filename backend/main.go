package main

import (
	"log"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/handlers"
	"github.com/vert3xc/barhat_tyagi/backend/middleware"
)

func main() {
	go http.HandleFunc("/register", handlers.Register)
	go http.HandleFunc("/login", handlers.Login)
	go http.HandleFunc("/", middleware.SessionHandler(handlers.Index))
	go http.HandleFunc("GET /logout", handlers.Logout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
