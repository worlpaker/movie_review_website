package Watchlist

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
	*Models.Watchlist | *Models.DeleteWatchlistMany
}

func ConvertJsontoBuf[T ModelsGeneric](t *testing.T, data T) *bytes.Buffer {
	databuf := new(bytes.Buffer)
	if err := json.NewEncoder(databuf).Encode(data); err != nil {
		t.Errorf("error in convert json to buf %v", err)
	}
	return databuf
}

func TestAddWatchlist(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/add_watchlist", AddWatchlist)
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "test"}
	databuf := ConvertJsontoBuf(t, add_watchlist)
	req, _ := http.NewRequest("POST", "/api/add_watchlist", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_data := &Models.DeleteDataModel{CollName: "watchlist", Filter: "movie_name", Data: add_watchlist.Movie_name}
	err := DB.Mongo.DeleteData(remove_data)
	assert.Empty(t, err)
}
