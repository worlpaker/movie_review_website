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

func TestShowAllReviews(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Get("/api/show_all_reviews", ShowAllReviews)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = DB.Mongo.AddReview(add_review)
	assert.Empty(t, err)
	req, _ := http.NewRequest("GET", "/api/show_all_reviews", nil)
	resp := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, resp.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = DB.Mongo.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowReviewsByEmail(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/review_by_email", ShowReviewsByEmail)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = DB.Mongo.AddReview(add_review)
	assert.Empty(t, err)
	databuf := ConvertJsontoBuf(t, add_review)
	req, _ := http.NewRequest("POST", "/api/review_by_email", databuf)
	resp := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, resp.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = DB.Mongo.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowReviewsByMovieName(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/show_reviews_by_movie", ShowReviewsByMovieName)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = DB.Mongo.AddReview(add_review)
	assert.Empty(t, err)
	databuf := ConvertJsontoBuf(t, add_review)
	req, _ := http.NewRequest("POST", "/api/show_reviews_by_movie", databuf)
	resp := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, resp.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = DB.Mongo.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestShowReviewsBy_Email_and_MovieName(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/review_movie_email", ShowReviewsBy_Email_and_MovieName)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = DB.Mongo.AddReview(add_review)
	assert.Empty(t, err)
	databuf := ConvertJsontoBuf(t, add_review)
	req, _ := http.NewRequest("POST", "/api/review_movie_email", databuf)
	resp := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, resp.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = DB.Mongo.DeleteData(remove_review)
	assert.Empty(t, err)
}

func TestCount_Reviews_and_Watchlist_ByEmail(t *testing.T) {
	log.SetOutput(io.Discard)
	if err := MongoDB.Mongoclient(); err != nil {
		t.Errorf("error in db connection %v", err)
	}
	router := chi.NewRouter()
	router.Post("/api/count_reviews_by_email", Count_Reviews_and_Watchlist_ByEmail)
	add_movie := &Models.Movies_model{Email: "test@test.com", Movie_name: "test", Movie_type: "test", Movie_cat: "test", Movie_year: 1900}
	err := DB.Mongo.AddMovie(add_movie)
	assert.Empty(t, err)
	add_review := &Models.Review_model{Email: "test@test.com", Movie_name: "test", Movie_rate: 10, Movie_review: "test"}
	err = DB.Mongo.AddReview(add_review)
	assert.Empty(t, err)
	add_watchlist := &Models.Watchlist{Email: "test@test.com", Movie_name: "testwatch"}
	err = DB.Mongo.AddWatchlist(add_watchlist)
	assert.Empty(t, err)
	databuf := ConvertJsontoBuf(t, add_review)
	req, _ := http.NewRequest("POST", "/api/count_reviews_by_email", databuf)
	resp := ExecuteRequest(req, router)
	assert.Equal(t, http.StatusOK, resp.Code)
	remove_movie := &Models.DeleteDataModel{CollName: "movies", Filter: "movie_name", Data: add_movie.Movie_name}
	err = DB.Mongo.DeleteData(remove_movie)
	assert.Empty(t, err)
	remove_review := &Models.DeleteDataModel{CollName: "review", Filter: "email", Data: add_review.Email}
	err = DB.Mongo.DeleteData(remove_review)
	assert.Empty(t, err)
	remove_watchlist := &Models.DeleteDataModel{CollName: "watchlist", Filter: "movie_name", Data: add_watchlist.Movie_name}
	err = DB.Mongo.DeleteData(remove_watchlist)
	assert.Empty(t, err)
}
