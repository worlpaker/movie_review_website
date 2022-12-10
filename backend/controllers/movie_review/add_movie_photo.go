package Review

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"
	"fmt"
	"net/http"
)

// AddPhotoMovie adds photo for movie
func AddPhotoMovie(w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(10 << 20)
	file, _, err := r.FormFile("movie_pic")
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	defer file.Close()
	movie_id := r.FormValue("movie_id")
	data := fmt.Sprintf("images/movies/%s.jpg", movie_id)
	if err := DB.Mongo.UploadPhoto(file, data); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: "success"}
	render.RenderJSON(w, r)
}
