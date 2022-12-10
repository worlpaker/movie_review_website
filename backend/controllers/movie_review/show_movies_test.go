package Review

import (
	MongoDB "backend/database/mongo"
	Models "backend/models"
	"io"
	"log"
	"net/http"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func TestShowAllMovies(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Get("/api/show_all_movies", ShowAllMovies)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	req, _ := http.NewRequest("GET", "/api/show_all_movies", nil)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
}

func TestShowMoviesByCat(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/movies_by_cat", ShowMoviesByCat)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	databuf := ConvertJsontoBuf(t, add_movie)
	req, _ := http.NewRequest("POST", "/api/movies_by_cat", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
}

func TestSearchMoviesByName(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/search_movies", SearchMoviesByName)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	databuf := ConvertJsontoBuf(t, add_movie)
	req, _ := http.NewRequest("POST", "/api/search_movies", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
}
