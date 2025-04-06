package handlers

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/vert3xc/barhat_tyagi/backend/utils"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

func ThreadVotings(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	threadID, err := strconv.Atoi(params["threadId"])
	if err != nil {
		http.Error(w, "Invalid thread ID format", http.StatusBadRequest)
		return
	}
	db, err := utils.ConnectToDb()
	if err != nil {
		http.Error(w, "Database connection error", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query(`
        SELECT id, title, descr 
        FROM votings 
        WHERE thread_id = $1
    `, threadID)

	if err != nil {
		http.Error(w, "Database query error", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	votings := make([]utils.Voting, 0)
	var id int
	var title, descr string
	for rows.Next() {
		if err := rows.Scan(&id, &title, &descr); err != nil {
			http.Error(w, "An error occurred with the database", http.StatusInternalServerError)
			return
		}
		voting := utils.Voting{
			ID:       id,
			ThreadId: threadID,
			Title:    title,
			Descr:    descr,
		}
		votings = append(votings, voting)
	}
	tmpl, err := template.ParseFiles("templates/threadVotings.html")
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, votings)
}
