package MongoDB

import (
	Models "backend/models"
	"fmt"
	"io"
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddMovie(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	create_data := []struct {
		data            *Models.Movies_model
		expected_output error
	}{
		{data: &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900},
			expected_output: nil},
		{data: &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900},
			expected_output: fmt.Errorf("movie already exist")},
	}
	for i, k := range create_data {
		name := fmt.Sprintf("Testno:%d", i)
		t.Run(name, func(t *testing.T) {
			err := Handle.AddMovie(k.data)
			assert.Equal(t, k.expected_output, err)
		})
	}
	remove_data := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: create_data[0].data.Movie_name}
	err := Handle.DeleteData(remove_data)
	assert.Empty(t, err)
}

func TestAddReview(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestEditMovie(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	edit_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test1", Movie_cat: "test1", Movie_year: 1900}
	err = Handle.EditMovie(edit_movie)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
}

func TestEditReview(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	edit_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 5, Movie_review: "test1"}
	err = Handle.EditReview(edit_review)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowAllMovies(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	_, err = Handle.ShowAllMovies()
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
}

func TestShowMoviesByCat(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	data := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(data)
	assert.Empty(t, err)
	_, err = Handle.ShowMoviesByCat(data)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: data.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
}

func TestSearchMoviesByName(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	data := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(data)
	assert.Empty(t, err)
	_, err = Handle.SearchMoviesByName(data)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: data.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
}

func TestShowAllReviews(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	_, err = Handle.ShowAllReviews()
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowReviewsByEmail(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	_, err = Handle.ShowReviewsByEmail(add_review)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowReviewsByMovieName(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	_, err = Handle.ShowReviewsByMovieName(add_review)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowReviewsBy_Email_and_MovieName(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	_, err = Handle.ShowReviewsBy_Email_and_MovieName(add_review)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestCount_Reviews_and_Watchlist_ByEmail(t *testing.T) {
	log.SetOutput(io.Discard)
	if db := Mongoclient(); db != nil {
		t.Errorf("error in db connection %v", db)
	}
	Handle := NewAuth()
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := Handle.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = Handle.AddReview(add_review)
	assert.Empty(t, err)
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "test"}
	err = Handle.AddWatchlist(add_watchlist)
	assert.Empty(t, err)
	user := &Models.Account{Email: "test@test.com"}
	_, err = Handle.Count_Reviews_and_Watchlist_ByEmail(user)
	assert.Empty(t, err)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = Handle.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = Handle.DeleteData(remove_review)
	assert.Empty(t, err)
	remove_watchlist := &Models.DeleteDataModel{CollName: "watchlist", Filter: "movie_name", Data: add_watchlist.Movie_name}
	err = Handle.DeleteData(remove_watchlist)
	assert.Empty(t, err)
}
