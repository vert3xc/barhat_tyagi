package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/vert3xc/barhat_tyagi/backend/handlers"
	"github.com/vert3xc/barhat_tyagi/backend/middleware"
)

func main() {
	r := mux.NewRouter()
	r.PathPrefix("/static/").Handler(
		http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))),
	)

	r.HandleFunc("/register", handlers.Register)
	r.HandleFunc("/login", handlers.Login)
	r.HandleFunc("/", middleware.SessionHandler(handlers.Index))
	r.HandleFunc("/create_voting", middleware.RoleMiddleware([]string{"Moderator", "Admin"})(handlers.CreateVoting))
	r.HandleFunc("/create_thread", middleware.RoleMiddleware([]string{"Moderator", "Admin"})(handlers.CreateThread))
	r.HandleFunc("/threads", middleware.SessionHandler(handlers.Threads))
	r.HandleFunc("/threads/{threadId}", middleware.SessionHandler(handlers.ThreadVotings))
	r.HandleFunc("/threads/{threadId}/{votingId}", middleware.SessionHandler(handlers.Index))
	r.HandleFunc("/logout", handlers.Logout)
	log.Fatal(http.ListenAndServe(":8080", r))
}
