package User

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"
	Models "backend/models"
	"encoding/json"
	"net/http"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user *Models.Account
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&user); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	data_token, err := DB.Mongo.Login(user)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	ts, err := DB.Redis.CreateToken(data_token)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	cookie := &http.Cookie{
		Name:     "Token",
		Value:    ts.RefreshToken,
		HttpOnly: false,
		MaxAge:   int(time.Hour * 24 * 3),
		Path:     "/",
	}
	http.SetCookie(w, cookie)
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: ts.AccessToken}
	render.Render(w, r)
}
