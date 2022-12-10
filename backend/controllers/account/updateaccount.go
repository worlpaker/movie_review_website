package User

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"
	Models "backend/models"
	"encoding/json"
	"net/http"
	"time"
)

// ChangePassword changes password of user
func ChangePassword(w http.ResponseWriter, r *http.Request) {
	var user *Models.ChangePassword
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&user); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	if err := DB.Mongo.ChangePassword(user); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: "password successfully changed"}
	render.RenderJSON(w, r)
}

// UpdateProfile updates informations of users
func UpdateProfile(w http.ResponseWriter, r *http.Request) {
	var user *Models.UpdateAccount
	if err := json.NewDecoder(r.Body).Decode(&user); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	new_token, err := DB.Mongo.UpdateProfile(user)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	ts, err := DB.Redis.CreateToken(new_token)
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
