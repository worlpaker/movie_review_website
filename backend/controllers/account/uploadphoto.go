package User

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"
	"fmt"
	"net/http"
)

func UploadPhoto(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20) // Parse our multipart form, 10 << 20 specifies a maximum
	file, _, err := r.FormFile("myFile")
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	defer file.Close()
	user_profile := r.FormValue("user_profile")
	data := fmt.Sprintf("images/profiles/%s.jpg", user_profile)
	if err := DB.Mongo.UploadPhoto(file, data); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: "successfully uploaded photo"}
	render.RenderJSON(w, r)
}
