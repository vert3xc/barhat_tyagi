package handlers

import (
	"html"
	"html/template"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/utils"

	_ "github.com/lib/pq"
)

func CreateVoting(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := utils.ConnectToDb()
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()
		threadName := r.FormValue("threadName")
		title := r.FormValue("title")
		descr := r.FormValue("descr")
		sanitizedTitle := html.EscapeString(title)
		sanitizedDescr := html.EscapeString(descr)
		if sanitizedTitle != title || sanitizedDescr != descr {
			http.Error(w, "Title or description contains ivalid characters", http.StatusBadRequest)
			return
		}
		options := r.Form["option"]
		if len(options) < 2 {
			http.Error(w, "At least 2 options required", http.StatusBadRequest)
			return
		}
		for _, option := range options {
			if option != html.EscapeString(option) {
				http.Error(w, "One of the options contains invalid characters", http.StatusBadRequest)
				return
			}
		}
		tx, err := db.Begin()
		defer tx.Rollback()
		var threadId int
		err = tx.QueryRow("SELECT id FROM threads WHERE thread_name = $1", threadName).Scan(&threadId)
		if err != nil {
			http.Error(w, "Invalid thread", http.StatusBadRequest)
			return
		}
		var votingId int
		err = tx.QueryRow(
			"INSERT INTO votings (thread_id, title, descr) VALUES ($1, $2, $3) RETURNING id",
			threadId, title, descr,
		).Scan(&votingId)
		if err != nil {
			http.Error(w, "Failed to create voting", http.StatusInternalServerError)
			return
		}
		for _, option := range options {
			sanitizedOption := html.EscapeString(option)
			if option != sanitizedOption {
				http.Error(w, "Invalid option content", http.StatusBadRequest)
				return
			}

			_, err = tx.Exec(
				"INSERT INTO options (voting_id, option_text) VALUES ($1, $2)",
				votingId, option,
			)
			if err != nil {
				http.Error(w, "Failed to save options", http.StatusInternalServerError)
				return
			}
		}
		err = tx.Commit()
		if err != nil {
			http.Error(w, "Failed to finalize voting", http.StatusInternalServerError)
			return
		}
	} else {
		tmpl, err := template.ParseFiles("templates/createVoting.html")
		w.Header().Set("Content-Type", "text/html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}

func CreateThread(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := utils.ConnectToDb()
		if err != nil {
			http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
			return
		}
		defer db.Close()
		threadName := r.FormValue("threadName")
		sanitizedThread := html.EscapeString(threadName)
		if sanitizedThread != threadName {
			http.Error(w, "Thread name contains invalid characters.", http.StatusBadRequest)
		}
		_, err = db.Exec(
			"INSERT INTO threads (thread_name) VALUES ($1)",
			threadName,
		)
	} else {
		tmpl, err := template.ParseFiles("templates/createThread.html")
		w.Header().Set("Content-Type", "text/html")
		if err != nil {
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, nil)
	}
}
