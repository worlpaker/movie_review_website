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

func TestEditMovie(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/edit_movie", EditMovie)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	edit_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test1", Movie_cat: "test1", Movie_year: 1900}
	databuf := ConvertJsontoBuf(t, edit_movie)
	req, _ := http.NewRequest("POST", "/api/edit_movie", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
}
