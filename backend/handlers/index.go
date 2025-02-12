package handlers

import (
	"html/template"
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func Index(w http.ResponseWriter, r *http.Request) {
	var contextKey utils.ContextKey = "session"
	sessionData, ok := r.Context().Value(contextKey).(utils.SessionData)
	if !ok {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
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
