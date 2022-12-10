package MongoDB

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

// InitAccountBSON initialize Account values before created
func InitAccountBSON(user *Models.Account) (*Models.Account, error) {
	if user.CreatedAt.IsZero() {
		user.CreatedAt = time.Now()
	}
	if user.UpdatedAt.IsZero() {
		user.UpdatedAt = time.Now()
	}
	user.Id = uuid.New()
	user.Profile_picture = uuid.New()
	user.Authorization = "user"
	return user, nil
}

// InitMovieBSON initialize Movie values before created
func InitMovieBSON(movie *Models.Movies_model) (*Models.Movies_model, error) {
	if movie.CreatedAt.IsZero() {
		movie.CreatedAt = time.Now()
	}
	if movie.UpdatedAt.IsZero() {
		movie.UpdatedAt = time.Now()
	}
	movie.Id = uuid.New()
	return movie, nil
}

// InitReviewBSON initialize Review values before created
func InitReviewBSON(review *Models.Review_model) (*Models.Review_model, error) {
	if review.CreatedAt.IsZero() {
		review.CreatedAt = time.Now()
	}
	if review.UpdatedAt.IsZero() {
		review.UpdatedAt = time.Now()
	}
	review.Id = uuid.New()
	return review, nil
}

// MarshalBSON, initialize values before created
func InitWatchlistBSON(watchlist *Models.Watchlist) (*Models.Watchlist, error) {
	if watchlist.CreatedAt.IsZero() {
		watchlist.CreatedAt = time.Now()
	}
	if watchlist.UpdatedAt.IsZero() {
		watchlist.UpdatedAt = time.Now()
	}
	watchlist.Id = uuid.New()
	watchlist.Watched = false
	return watchlist, nil
}

// HashPassword hashes the password
func HashPassword(pwd []byte) string {
	hash, err := bcrypt.GenerateFromPassword(pwd, bcrypt.MinCost)
	if ErrHandler.Log(err) {
		return ""
	}
	return string(hash)
}

// VerifyPassword for user login
func VerifyPassword(dbPass, userPass string) error {
	if err := bcrypt.CompareHashAndPassword([]byte(dbPass), []byte(userPass)); ErrHandler.Log(err) {
		return err
	}
	return nil
}
