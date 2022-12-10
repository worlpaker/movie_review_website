package Models

type Movies_model struct {
	Base
	Email      string `json:"Email"`
	Movie_name string `json:"Movie_name"`
	Movie_type string `json:"Movie_type"`
	Movie_cat  string `json:"Movie_cat"`
	Movie_year int    `json:"Movie_year"`
	Movie_rate int    `json:"Movie_rate"`
}

type Picture_model struct {
	Movie_id string `json:"movie_id"`
}

type Review_model struct {
	Base
	Email         string `json:"Email"`
	Movie_name    string `json:"Movie_name"`
	Movie_review  string `json:"Movie_review"`
	Movie_rate    int    `json:"Movie_rate"`
	Movie_picture string `json:"Movie_picture"`
}
