package middleware

import (
	"context"
	"net/http"
	"time"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func SessionHandler(f http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("session")
		var sessionData utils.SessionData
		if err != nil {
			sessionData := utils.SessionData{
				ID:       0,
				Username: "Anon",
                                Role: "Anon",
				Expiry:   time.Now().Add(24 * time.Hour),
			}
			session, err := utils.CreateSession(sessionData)
			if err != nil {
				return
			}
			http.SetCookie(w, &http.Cookie{
				Name:     "session",
				Value:    session,
				Expires:  sessionData.Expiry,
				HttpOnly: true,
				Secure:   false,
				SameSite: http.SameSiteLaxMode,
			})
		} else {
			sessionData, err = utils.DecodeSession(cookie.Value)
			if err != nil {
				ClearSessionCookie(w)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}
		}
		var contextKey utils.ContextKey = "session"
		ctx := context.WithValue(r.Context(), contextKey, sessionData)
		f.ServeHTTP(w, r.WithContext(ctx))
	})
}

func ClearSessionCookie(w http.ResponseWriter) {
	http.SetCookie(w, &http.Cookie{
		Name:     "session",
		Value:    "",
		MaxAge:   -1,
		HttpOnly: true,
		Secure:   false,
		SameSite: http.SameSiteLaxMode,
	})
}
