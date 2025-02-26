package main

import (
	"log"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/handlers"
	"github.com/vert3xc/barhat_tyagi/backend/middleware"
)

func main() {
	http.HandleFunc("/register", handlers.Register)
	http.HandleFunc("/login", handlers.Login)
	http.HandleFunc("/", middleware.SessionHandler(handlers.Index))
	http.HandleFunc("/logout", handlers.Logout)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
