package handlers

import (
	"net/http"

	"github.com/vert3xc/barhat_tyagi/backend/middleware"
)

func Logout(w http.ResponseWriter, r *http.Request) {
	middleware.ClearSessionCookie(w)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
