package middleware

import (
	"log"
	"net/http"
	"slices"

	"github.com/vert3xc/barhat_tyagi/backend/utils"
)

func RoleMiddleware(allowedRoles []string) func(http.HandlerFunc) http.HandlerFunc {
	return func(next http.HandlerFunc) http.HandlerFunc {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			cookie, err := r.Cookie("session")
			if err != nil {
				ClearSessionCookie(w)
				http.Error(w, "Unauthorized: No session found", http.StatusUnauthorized)
				return
			}
			sessionData, err := utils.DecodeSession(cookie.Value)
			if err != nil {
				log.Printf("Session decoding error: %v", err)
				ClearSessionCookie(w)
				http.Error(w, "Invalid session", http.StatusUnauthorized)
				return
			}
			if !slices.Contains(allowedRoles, sessionData.Role) {
				http.Error(w, "Insufficient priveleges", http.StatusForbidden)
				return
			}
			next.ServeHTTP(w, r)
		})
	}
}
