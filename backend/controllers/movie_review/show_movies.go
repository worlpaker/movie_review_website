package Review

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"encoding/json"
	"net/http"
)

// ShowAllMovies deploys movies to frontend
func ShowAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := DB.Mongo.ShowAllMovies()
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(movies)
}

// ShowMoviesByCat show movies by category
func ShowMoviesByCat(w http.ResponseWriter, r *http.Request) {
	var movie *Models.Movies_model
	if err := json.NewDecoder(r.Body).Decode(&movie); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	filtered, err := DB.Mongo.ShowMoviesByCat(movie)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)

}

// SearchMoviesByName handles search movie by name
func SearchMoviesByName(w http.ResponseWriter, r *http.Request) {
	var movie *Models.Movies_model
	if err := json.NewDecoder(r.Body).Decode(&movie); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	filtered, err := DB.Mongo.SearchMoviesByName(movie)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)
}
