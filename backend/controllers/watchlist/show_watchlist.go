package Watchlist

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"encoding/json"
	"net/http"
)

// ShowWatchlistByEmail shows watchlist movies by email
func ShowWatchlistByEmail(w http.ResponseWriter, r *http.Request) {
	var watchlist *Models.Watchlist
	if err := json.NewDecoder(r.Body).Decode(&watchlist); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	filtered, err := DB.Mongo.ShowWatchlistByEmail(watchlist)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)
}
