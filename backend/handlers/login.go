package handlers

import (
	"crypto/sha256"
	"encoding/hex"
	"html"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/vert3xc/barhat_tyagi/backend/utils"

	_ "github.com/lib/pq"
)

func Login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		db, err := utils.ConnectToDb()
		if err != nil {
			http.Error(w, "Database connection error", http.StatusInternalServerError)
			return
		}
		defer db.Close()
		username := r.FormValue("username")
		password := r.FormValue("passwd")
		if username == "" || password == "" {
			http.Error(w, "Username and password are required", http.StatusBadRequest)
			return
		}
		sanitizedUsername := html.EscapeString(username)
		log.Println(sanitizedUsername)
		if sanitizedUsername != username {
			http.Error(w, "Username contains invalid characters.", http.StatusBadRequest)
			return
		}
		hash := sha256.Sum256([]byte(password))
		hshdPassword := hex.EncodeToString(hash[:])
		var id int
		var actual_psswd string
		err = db.QueryRow(
			"SELECT id, password_hash FROM users WHERE username = $1",
			username,
		).Scan(&id, &actual_psswd)
		if err != nil {
			http.Error(w, "problem with finding user", http.StatusInternalServerError)
			return
		}
		if hshdPassword != actual_psswd {
			http.Error(w, "invalid username or password", http.StatusUnauthorized)
			return
		}
		sessionData := utils.SessionData{
			ID:       id,
			Username: username,
			Expiry:   time.Now().Add(24 * time.Hour),
		}
		session, err := utils.CreateSession(sessionData)
		if err != nil {
			http.Error(w, "Problem creating session", http.StatusUnauthorized)
		}
		http.SetCookie(w, &http.Cookie{
			Name:     "session",
			Value:    session,
			Expires:  sessionData.Expiry,
			HttpOnly: true,
			Secure:   false,
			SameSite: http.SameSiteLaxMode,
		})
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl, err := template.ParseFiles("templates/login.html")
	w.Header().Set("Content-Type", "text/html")
	if err != nil {
		http.Error(w, "Template not found", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}
