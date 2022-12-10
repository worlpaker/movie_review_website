package Review

import (
	MongoDB "backend/database/mongo"
	Models "backend/models"
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/stretchr/testify/assert"
)

func ExecuteRequest(r *http.Request, s *chi.Mux) *httptest.ResponseRecorder {
	w := httptest.NewRecorder()
	s.ServeHTTP(w, r)
	return w
}

type ModelsGeneric interface {
	*Models.Movies_model | *Models.Review_model
}

func ConvertJsontoBuf[T ModelsGeneric](t *testing.T, data T) *bytes.Buffer {
	databuf := new(bytes.Buffer)
	if err := json.NewEncoder(databuf).Encode(data); err != nil {
		t.Errorf("error in convert json to buf %v", err)
	}
	return databuf
}
func TestAddMovie(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/add_movie", AddMovie)
	data := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	databuf := ConvertJsontoBuf(t, data)
	req, _ := http.NewRequest("POST", "/api/add_movie", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_data := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: data.Movie_name}
	err := DB.Mongo.DeleteData(remove_data)
	assert.Empty(t, err)

}
