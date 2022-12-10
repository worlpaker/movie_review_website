package User

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"
	"net/http"
	"time"
)

// Logout removes old cookie
func Logout(w http.ResponseWriter, r *http.Request) {
	token, err := r.Cookie("Token")
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	if err := DB.Redis.CheckToken(token.Value); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	if err := DB.Redis.DeleteToken(token.Value); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    "",
		HttpOnly: false,
		Path:     "/",
		Expires:  time.Unix(0, 0),
	}
	http.SetCookie(w, cookie)
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: "Successfully logged out"}
	render.RenderJSON(w, r)
}
