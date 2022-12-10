package MongoDB

import (
	Models "backend/models"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddWatchlist(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	create_data := []struct {
		data            *Models.Watchlist
		expected_output error
	}{
		{data: &Models.Watchlist{Email: "test@test.com", Movie_name: "test"},
			expected_output: nil},
		{data: &Models.Watchlist{Email: "test@test.com", Movie_name: "test"},
			expected_output: fmt.Errorf("watchlist movie already exist")},
	}
	for i, k := range create_data {
		name := fmt.Sprintf("Testno:%d", i)
		t.Run(name, func(t *testing.T) {
			err := Handle.AddWatchlist(k.data)
			assert.Equal(t, k.expected_output, err)
		})
	}
	remove_data := &Models.DeleteDataModel{CollName: "watchlist", Filter: "movie_name", Data: create_data[0].data.Movie_name}
	err := Handle.DeleteData(remove_data)
	assert.Empty(t, err)
}

func TestDeleteWatchlist(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "test"}
	err := Handle.AddWatchlist(add_watchlist)
	assert.Empty(t, err)
	edit_watchlist := &Models.DeleteWatchlistMany{Email: "test@test.com", Movie_name: []string{"test"}}
	err = Handle.DeleteWatchlist(edit_watchlist)
	assert.Empty(t, err)
}

func TestShowWatchlistByEmail(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "test"}
	err := Handle.AddWatchlist(add_watchlist)
	assert.Empty(t, err)
	_, err = Handle.ShowWatchlistByEmail(add_watchlist)
	assert.Empty(t, err)
	remove_watchlist := &Models.DeleteDataModel{CollName: "watchlist", Filter: "movie_name", Data: add_watchlist.Movie_name}
	err = Handle.DeleteData(remove_watchlist)
	assert.Empty(t, err)
}
