package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"html"
	"html/template"
	"log"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/utils"

	_ "github.com/lib/pq"
)

func Register(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := utils.ConnectToDb()
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()
		username := r.FormValue("username")
		log.Println(username)
		sanitizedUsername := html.EscapeString(username)
		log.Println(sanitizedUsername)
		if sanitizedUsername != username {
			http.Error(w, "Username contains invalid characters.", http.StatusBadRequest)
			return
		}
		password := r.FormValue("passwd")
		log.Println(password)
		if username == "" || password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}
		hash := sha256.Sum256([]byte(password))
		hshdPassword := hex.EncodeToString(hash[:])
		_, err = db.Exec(
			"INSERT INTO users (username, password_hash) VALUES ($1, $2)",
			sanitizedUsername, hshdPassword,
		)

		if err != nil {
			http.Error(w, "Problem adding user. Probably username already exists", http.StatusConflict)
			return
		}
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("templates/register.html")
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
