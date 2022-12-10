package Watchlist

import (
	ErrHandler "backend/helper_handlers/error"
	SuccessHandler "backend/helper_handlers/success"
	Models "backend/models"
	"encoding/json"
	"net/http"
)

// AddWatchlist, adds a new movie to watch list
func AddWatchlist(w http.ResponseWriter, r *http.Request) {
	var watchlist *Models.Watchlist
	decoder := json.NewDecoder(r.Body)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(&watchlist); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	if err := DB.Mongo.AddWatchlist(watchlist); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	render := &SuccessHandler.Response{HTTPStatusCode: http.StatusOK, Message: "successfully add a movie to watchlist"}
	render.RenderJSON(w, r)
}
