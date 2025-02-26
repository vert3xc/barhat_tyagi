package handlers

import (
	"html/template"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	var contextKey utils.ContextKey = "session"
	sessionData, ok := r.Context().Value(contextKey).(utils.SessionData)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	name := template.HTML(sessionData.Username)
	if name == "" {
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	tmpl, err := template.New("greeting").Parse(`<b>Hello, {{.}}</b>`)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, name)
}
