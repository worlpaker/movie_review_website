package Review

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"encoding/json"
	"net/http"
)

// ShowAllReviews deploys reviews to frontend from latest
func ShowAllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := DB.Mongo.ShowAllReviews()
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(reviews)
}

// ShowReviewsByEmail show all reviews by user email
func ShowReviewsByEmail(w http.ResponseWriter, r *http.Request) {
	var review *Models.Review_model
	if err := json.NewDecoder(r.Body).Decode(&review); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	filtered, err := DB.Mongo.ShowReviewsByEmail(review)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)
}

// ShowReviewsByEmail show all reviews by movie name
func ShowReviewsByMovieName(w http.ResponseWriter, r *http.Request) {
	var review *Models.Review_model
	if err := json.NewDecoder(r.Body).Decode(&review); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	filtered, err := DB.Mongo.ShowReviewsByMovieName(review)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)
}

// ShowReviewsBy_Email_and_MovieName show all reviews by user email regarding to movie name
func ShowReviewsBy_Email_and_MovieName(w http.ResponseWriter, r *http.Request) {
	var review *Models.Review_model
	if err := json.NewDecoder(r.Body).Decode(&review); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	filtered, err := DB.Mongo.ShowReviewsBy_Email_and_MovieName(review)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(filtered)
}

// Count_Reviews_and_Watchlist_ByEmail, counts reviews and watchlist by user email
func Count_Reviews_and_Watchlist_ByEmail(w http.ResponseWriter, r *http.Request) {
	var review *Models.Account
	if err := json.NewDecoder(r.Body).Decode(&review); ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	result, err := DB.Mongo.Count_Reviews_and_Watchlist_ByEmail(review)
	if ErrHandler.Log(err) {
		render := &ErrHandler.Response{HTTPStatusCode: http.StatusBadRequest, StatusText: http.StatusText(http.StatusBadRequest), ErrorText: err}
		render.ErrRequest(w, r)
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}
