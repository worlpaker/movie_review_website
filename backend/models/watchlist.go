package Models

type Watchlist struct {
	Base
	Email      string `json:"Email"`
	Movie_name string `json:"Movie_name"`
	Watched    bool   `json:"Watched"`
}

type DeleteWatchlistMany struct {
	Email      string   `json:"Email"`
	Movie_name []string `json:"Movie_name"`
}
