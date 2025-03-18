package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var contextKey utils.ContextKey = "session"
	sessionData, ok := r.Context().Value(contextKey).(utils.SessionData)
	if !ok || sessionData.Username == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
    
	tmplPath := filepath.Join("templates", "index.html")
	tmpl, err := template.ParseFiles(tmplPath)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	err = tmpl.Execute(w, struct {
		Username string
	}{
		Username: sessionData.Username,
	})
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
