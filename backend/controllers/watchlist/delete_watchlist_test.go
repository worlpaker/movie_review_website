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

func TestDeleteWatchlist(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/delete_watchlist", DeleteWatchlist)
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "test"}
	err := DB.Mongo.AddWatchlist(add_watchlist)
	assert.Empty(t, err)
	edit_watchlist := &Models.DeleteWatchlistMany{Email: "test@test.com", Movie_name: []string{"test"}}
	databuf := ConvertJsontoBuf(t, edit_watchlist)
	req, _ := http.NewRequest("POST", "/api/delete_watchlist", databuf)
	response := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, response.Code)
}
