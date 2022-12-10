package Server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"

	User "backend/controllers/account"
	Review "backend/controllers/movie_review"
	Watchlist "backend/controllers/watchlist"
)

// NewRouter creates a new route
func NewRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	// Basic CORS
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	fs := http.FileServer(http.Dir("images"))
	r.Handle("/api/images/*", http.StripPrefix("/api/images", fs))
	return r
}

// userRouter for /api/profile/
func userRouter() chi.Router {
	r := chi.NewRouter()
	r.Use(UsersOnly)
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("User Page"))
	})
	r.Post("/logout", User.Logout)
	r.Post("/refresh", User.Refresh)
	r.Post("/changepassword", User.ChangePassword)
	r.Put("/updateprofile", User.UpdateProfile)
	return r
}

// SetupRouters for api
func SetupRouters() *chi.Mux {
	r := NewRouter()
	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("api"))
	})

	r.Post("/api/register", User.CreateUser)
	r.Post("/api/uploadphoto", User.UploadPhoto)
	r.Post("/api/login", User.Login)

	r.Post("/api/add_movie", Review.AddMovie)
	r.Post("/api/add_movie_photo", Review.AddPhotoMovie)
	r.Post("/api/edit_movie", Review.EditMovie)
	r.Post("/api/movies_by_cat", Review.ShowMoviesByCat)
	r.Post("/api/search_movies", Review.SearchMoviesByName)
	r.Get("/api/show_all_movies", Review.ShowAllMovies)
	r.Post("/api/add_review", Review.AddReview)
	r.Post("/api/edit_review", Review.EditReview)
	r.Post("/api/show_reviews_by_movie", Review.ShowReviewsByMovieName)
	r.Post("/api/review_by_email", Review.ShowReviewsByEmail)
	r.Post("/api/count_reviews_by_email", Review.Count_Reviews_and_Watchlist_ByEmail)
	r.Post("/api/review_movie_email", Review.ShowReviewsBy_Email_and_MovieName)
	r.Get("/api/show_all_reviews", Review.ShowAllReviews)

	r.Post("/api/add_watchlist", Watchlist.AddWatchlist)
	r.Post("/api/delete_watchlist", Watchlist.DeleteWatchlist)
	r.Post("/api/show_watchlist", Watchlist.ShowWatchlistByEmail)

	r.Mount("/api/profile", userRouter())
	return r
}

// Start to server
func Start(port string) error {
	r := SetupRouters()
	return http.ListenAndServe(port, r)
}
