package handlers

import (
	"html/template"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/utils"

	_ "github.com/lib/pq"
)

func Threads(w http.ResponseWriter, r *http.Request) {
	db, err := utils.ConnectToDb()
	if err != nil {
		http.Error(w, "Problem connecting to database", http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT id, thread_name FROM threads")
	if err != nil {
		http.Error(w, "Problem connecting to database", http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	threads := make([]utils.Thread, 0)
	for rows.Next() {
		var threadName string
		var id int
		if err := rows.Scan(&id, &threadName); err != nil {
			http.Error(w, "A problem occurred with the database", http.StatusInternalServerError)
			return
		}
		thread := utils.Thread{
			ID:         id,
			ThreadName: threadName,
		}
		threads = append(threads, thread)
	}
	tmpl, err := template.ParseFiles("templates/threads.html")
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, threads)
}
