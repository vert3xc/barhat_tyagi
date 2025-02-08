package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

type ContextKey string

func SessionHandler(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		if err != nil {
			ClearSessionCookie(w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		sessionData, err := utils.DecodeSession(cookie.Value)
		if err != nil {
			ClearSessionCookie(w)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		var contextKey ContextKey = "session"
		ctx := context.WithValue(r.Context(), contextKey, sessionData)
		f.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		Expires:  time.Unix(0, 0),
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}
