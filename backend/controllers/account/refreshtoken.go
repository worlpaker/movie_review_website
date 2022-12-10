package User

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"

	"net/http"
)

// Refresh checks refresh token and create new access token
func Refresh(w http.ResponseWriter, r *http.Request) {
	refresh_token, err := DB.Redis.TokenValid(r)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	data, err := DB.Redis.ReadToken(refresh_token)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	ts, err := DB.Redis.CreateToken(data)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusUnauthorized, StatusText: http.StatusText(http.StatusUnauthorized), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: ts.AccessToken}
	render.Render(w, r)
}
