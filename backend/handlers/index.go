package handlers

import (
	"html/template"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	cookie, err := r.Cookie("session")
	if err != nil {
		http.Error(w, "no session found", http.StatusUnauthorized)
		return
	}
	sessionData, err := utils.DecodeSession(cookie.String())
	if err != nil {
		http.Error(w, "Invalid session", http.StatusUnauthorized)
		return
	}
	name := sessionData.Username
	tmpl, err := template.New("greeting").Parse(`<b>Hello, {{.}}</b>`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, name)
}
