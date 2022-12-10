package Watchlist

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

func TestShowWatchlistByEmail(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/show_watchlist", ShowWatchlistByEmail)
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "test"}
	databuf := ConvertJsontoBuf(t, add_watchlist)
	req, _ := http.NewRequest("POST", "/api/show_watchlist", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
	remove_data := &Models.DeleteDataModel{CollName: "watchlist", Filter: "movie_name", Data: add_watchlist.Movie_name}
	err := DB.Mongo.DeleteData(remove_data)
	assert.Empty(t, err)
}
