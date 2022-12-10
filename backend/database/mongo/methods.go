package MongoDB

import (
	Models "backend/models"
	"mime/multipart"
)

type service struct{}

func NewAuth() *service {
	return &service{}
}

type IHandlers interface {
	Login(*Models.Account) (*Models.Token, error)
	CreateUser(*Models.Account) (*Models.Account, error)
	DeleteData(*Models.DeleteDataModel) error
	ChangePassword(*Models.ChangePassword) error
	UpdateProfile(*Models.UpdateAccount) (*Models.Token, error)
	UploadPhoto(multipart.File, string) error
	AddMovie(*Models.Movies_model) error
	AddReview(*Models.Review_model) error
	EditMovie(*Models.Movies_model) error
	EditReview(*Models.Review_model) error
	ShowAllMovies() ([]Models.Movies_model, error)
	ShowMoviesByCat(*Models.Movies_model) ([]Models.Movies_model, error)
	SearchMoviesByName(*Models.Movies_model) ([]Models.Movies_model, error)
	ShowAllReviews() ([]Models.Review_model, error)
	ShowReviewsByEmail(*Models.Review_model) ([]Models.Review_model, error)
	ShowReviewsByMovieName(*Models.Review_model) ([]Models.Review_model, error)
	ShowReviewsBy_Email_and_MovieName(*Models.Review_model) ([]Models.Review_model, error)
	Count_Reviews_and_Watchlist_ByEmail(*Models.Account) (map[string]int64, error)
	AddWatchlist(*Models.Watchlist) error
	DeleteWatchlist(*Models.DeleteWatchlistMany) error
	ShowWatchlistByEmail(*Models.Watchlist) ([]Models.Watchlist, error)
}

var _ IHandlers = &service{}
