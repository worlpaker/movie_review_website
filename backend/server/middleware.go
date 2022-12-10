package Server

import (
	ErrHandler "backend/helper_handlers/error"
	"net/http"
)

// UsersOnly middleware to restricts access for users
func UsersOnly(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if _, err := r.Cookie("Token"); ErrHandler.Log(err) {
			render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
			render.ErrRequest(w, r)
			return
		}
		if _, err := DB.Redis.TokenValid(r); ErrHandler.Log(err) {
			render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
			render.ErrRequest(w, r)
			return
		}
		next.ServeHTTP(w, r)
	})
}
