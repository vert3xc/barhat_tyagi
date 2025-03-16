package handlers

import (
    "net/http"
    "github.com/vert3xc/barhat_tyagi/backend/utils"
    if "html"
    "fmt"

    _ "github.com/lib/pq"
)

func CreateVoting(w http.ResponseWriter, r *http.Request){
    if r.Method == "POST"{
        db, err := utils.ConnectToDb()
        if err != nil{
            http.Error(w, "Database connection error", http.StatusInternalServerError)
            return
        }
        defer db.Close()
        threadName := r.FormValue("threadName")
        title := r.FormValue("title")
        descr := r.FormValue("descr")
        sanitizedTitle := html.EcapeString(title)
        sanitizedDescr := html.EscapeString(descr)
        if title != sanitizedTitle || descr != sanitizedDescr {
            http.Error(w, "Title or description contains invalid characters.")
            return
        }
        var thread_id int
        err = db.QueryRow(
              "SELECT id FROM threads WHERE thread_name = $1",
              threadName,
        ).Scan(&thread_id)
        _, err = db.Exec(
                 "INSERT INTO votings (thread_id, title, descr) VALUES ($1, $2, $3)",
                 thread_id, title, descr,
        )
    } else{
        w.Header().Set("Content-Type", "text/html")
        fmt.Fprintf(w, "")
    }
}

func CreateThread(w http.ResponseWriter, r *http.Request){
    if r.Method == "POST"{
        db, err := utils.ConnectToDb()
        if err != nil{
            http.Error(w, "Failed to connect to database", http.StatusInternalServerError)
            return
        }
        defer db.Close()
        threadName := r.FormValue("threadName")
        sanitizedThread := html.EscapeString(threadName)
        if sanitizedThread != threadName{
            http.Error(w, "Thread name contains invalid characters.")
        }
        _, err = db.Exec(
            "INSERT INTO threads (thread_name) VALUES ($1)",
            threadName,
        )
    } else {
        fmt.Fprintf(w, "aboba %s", "sigma")
    }
}
