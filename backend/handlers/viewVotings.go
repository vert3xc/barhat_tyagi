package handlers

import (
	"html/template"
	"net/http"
	"strconv"
	"log"
	"github.com/gorilla/mux"
	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func ViewVotings(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		params := mux.Vars(r)
		votingId, err := strconv.Atoi(params["votingId"])
		if err != nil {
			http.Error(w, "Voting id is a non-integer value", http.StatusBadRequest)
			return
		}
		threadId, err := strconv.Atoi(params["threadId"])
		if err != nil {
			http.Error(w, "Thread id is a non-integer value", http.StatusBadRequest)
			return
		}
		db, err := utils.ConnectToDb()
		if err != nil {
			http.Error(w, "Problem connecting to database", http.StatusInternalServerError)
			return
		}
		defer db.Close()

		var title, descr string
		err = db.QueryRow("SELECT title, descr FROM votings WHERE id = $1", votingId).Scan(&title, &descr)
		if err != nil {
			http.Error(w, "Voting not found", http.StatusNotFound)
			return
		}

		voting := utils.Voting{
			ID:       votingId,
			ThreadId: threadId,
			Title:    title,
			Descr:    descr,
		}

		rows, err := db.Query("SELECT option_text, vote_count FROM options WHERE voting_id = $1", votingId)
		if err != nil {
			http.Error(w, "Problem retrieving options", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		options := make([]utils.Option, 0)

		var text string
		var count int
		for rows.Next() {
			if err := rows.Scan(&text, &count); err != nil {
				http.Error(w, "An error occurred with the database", http.StatusInternalServerError)
				return
			}
			option := utils.Option{
				VotingId:   votingId,
				OptionText: text,
				VoteCount:  count,
			}
			options = append(options, option)
		}

		// Check for errors from iterating over rows
		if err = rows.Err(); err != nil {
			http.Error(w, "Error processing options", http.StatusInternalServerError)
			return
		}

		tmpl, err := template.ParseFiles("templates/viewVoting.html")
		if err != nil {
			log.Println("Error parsing template:", err)
			http.Error(w, "Template not found", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/html")
		err = tmpl.Execute(w, utils.FullVoting{
			Voting:  voting,
			Options: options,
		})
		if err != nil {
			log.Println("Error executing template:", err)
		}
	}
}