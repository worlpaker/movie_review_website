package MongoDB

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (s *service) AddMovie(movies *Models.Movies_model) error {
	var dbmovies *Models.Movies_model
	collection := client.Database("GODB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ifexist := collection.FindOne(ctx, bson.M{"movie_name": movies.Movie_name}).Decode(&dbmovies); ifexist == nil {
		ErrHandler.Log(fmt.Errorf("movie already exist"))
		return fmt.Errorf("movie already exist")
	}
	movies, err := InitMovieBSON(movies)
	if ErrHandler.Log(err) {
		return err
	}
	if _, err := collection.InsertOne(ctx, movies); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) AddReview(review *Models.Review_model) error {
	var dbmoview *Models.Movies_model
	var results []bson.M
	collection_movies := client.Database("GODB").Collection("movies")
	ctx_movies, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection_movies.FindOne(ctx_movies, bson.M{"movie_name": review.Movie_name}).Decode(&dbmoview); ErrHandler.Log(err) {
		return err
	}
	collection_review := client.Database("GODB").Collection("review")
	ctx_review, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	review.Movie_picture = fmt.Sprint(dbmoview.Id)
	review, err := InitReviewBSON(review)
	if ErrHandler.Log(err) {
		return err
	}
	//Create new review
	if _, err := collection_review.InsertOne(ctx_review, review); ErrHandler.Log(err) {
		return err
	}
	//Average review
	pipeline := []bson.M{
		{
			"$match": bson.M{"movie_name": review.Movie_name}},
		{
			"$group": bson.M{
				"_id":     "",
				"average": bson.M{"$avg": "$movie_rate"},
			},
		},
	}
	resultcursor, err := collection_review.Aggregate(ctx_review, pipeline)
	if ErrHandler.Log(err) {
		return err
	}
	if err := resultcursor.All(context.TODO(), &results); ErrHandler.Log(err) {
		return err
	}
	average_value := results[0]["average"]
	//update movie_rate average.
	filter := bson.M{"movie_name": review.Movie_name}
	update := bson.M{"$set": bson.M{"movie_rate": average_value}}
	if _, err := collection_movies.UpdateOne(ctx_movies, filter, update); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) EditMovie(movies *Models.Movies_model) error {
	var dbmovies *Models.Movies_model
	collection := client.Database("GODB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection.FindOne(ctx, bson.M{"email": movies.Email, "movie_name": movies.Movie_name}).Decode(&dbmovies); ErrHandler.Log(err) {
		return err
	}
	filter := bson.M{"email": movies.Email, "movie_name": movies.Movie_name}
	update := bson.M{"$set": bson.M{"movie_type": movies.Movie_type, "movie_cat": movies.Movie_cat, "movie_year": movies.Movie_year,
		"base": bson.M{"createdat": dbmovies.CreatedAt, "updatedat": time.Now(), "deletedat": dbmovies.DeletedAt}}}

	if _, err := collection.UpdateOne(ctx, filter, update); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) EditReview(review *Models.Review_model) error {
	var dbreview *Models.Review_model
	var results []bson.M
	// part1: update review
	collection_review := client.Database("GODB").Collection("review")
	ctx_review, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := collection_review.FindOne(ctx_review, bson.M{"email": review.Email, "movie_name": review.Movie_name}).Decode(&dbreview); ErrHandler.Log(err) {
		return err
	}
	filter := bson.M{"email": review.Email, "movie_name": review.Movie_name}
	update := bson.M{"$set": bson.M{"movie_rate": review.Movie_rate, "movie_review": review.Movie_review,
		"base": bson.M{"createdat": dbreview.CreatedAt, "updatedat": time.Now(), "deletedat": dbreview.DeletedAt}}}

	if _, err := collection_review.UpdateOne(ctx_review, filter, update); ErrHandler.Log(err) {
		return err
	}
	//part2: update movie rate
	collection_movies := client.Database("GODB").Collection("movies")
	ctx_movies, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//Average review
	pipeline := []bson.M{
		{
			"$match": bson.M{"movie_name": review.Movie_name}},
		{
			"$group": bson.M{
				"_id":     "",
				"average": bson.M{"$avg": "$movie_rate"},
			},
		},
	}
	resultcursor, err := collection_review.Aggregate(ctx_review, pipeline)
	if ErrHandler.Log(err) {
		return err
	}
	if err := resultcursor.All(context.TODO(), &results); ErrHandler.Log(err) {
		return err
	}
	average_value := results[0]["average"]
	//update movie_rate average.
	filter = bson.M{"movie_name": review.Movie_name}
	update = bson.M{"$set": bson.M{"movie_rate": average_value}}
	if _, err := collection_movies.UpdateOne(ctx_movies, filter, update); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) ShowAllMovies() ([]Models.Movies_model, error) {
	var movies []Models.Movies_model
	collection := client.Database("GODB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	myOptions := options.Find()
	myOptions.SetSort(bson.M{"$natural": -1})
	cursor, err := collection.Find(ctx, bson.M{}, myOptions)
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err := cursor.All(ctx, &movies); ErrHandler.Log(err) {
		return nil, err
	}
	return movies, nil
}

func (s *service) ShowMoviesByCat(movie *Models.Movies_model) ([]Models.Movies_model, error) {
	var filtered []Models.Movies_model
	collection := client.Database("GODB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"movie_cat": movie.Movie_cat})
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err = cursor.All(ctx, &filtered); ErrHandler.Log(err) {
		return nil, err
	}
	return filtered, nil
}

// SearchMoviesByName, please see comments in the function.
func (s *service) SearchMoviesByName(movie *Models.Movies_model) ([]Models.Movies_model, error) {
	var filtered []Models.Movies_model
	collection := client.Database("GODB").Collection("movies")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// only for test, everytime you search movies it creates another index which is not useful.
	// instead that, before search a word, create text index as:
	// db.movies.createIndex({ "movie_name": "text" });
	// then search:
	// db.movies.find( { $text: { $search: "test" } } )
	// If you use above solution then remove optional part.
	// references:
	// [1] https://www.mongodb.com/docs/manual/core/index-text/
	// [2] https://www.mongodb.com/docs/drivers/go/v1.8/fundamentals/indexes/
	//optional start
	indexModel := mongo.IndexModel{
		Keys: bson.D{{Key: "movie_name", Value: "text"}},
	}
	if name, err := collection.Indexes().CreateOne(ctx, indexModel); ErrHandler.Log(err) {
		log.Println("NAME", name)
		return nil, err
	}
	//optional end
	cursor, err := collection.Find(ctx, bson.M{"$text": bson.M{"$search": movie.Movie_name}})
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err := cursor.All(ctx, &filtered); ErrHandler.Log(err) {
		return nil, err
	}
	return filtered, nil
}

func (s *service) ShowAllReviews() ([]Models.Review_model, error) {
	var reviews []Models.Review_model
	collection := client.Database("GODB").Collection("review")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	//sort from last review
	myOptions := options.Find()
	myOptions.SetSort(bson.M{"$natural": -1})
	cursor, err := collection.Find(ctx, bson.M{}, myOptions)
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err = cursor.All(ctx, &reviews); ErrHandler.Log(err) {
		return nil, err
	}
	return reviews, nil
}

func (s *service) ShowReviewsByEmail(review *Models.Review_model) ([]Models.Review_model, error) {
	var filtered []Models.Review_model
	collection := client.Database("GODB").Collection("review")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"email": review.Email})
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err = cursor.All(ctx, &filtered); ErrHandler.Log(err) {
		return nil, err
	}
	return filtered, nil
}

func (s *service) ShowReviewsByMovieName(review *Models.Review_model) ([]Models.Review_model, error) {
	var filtered []Models.Review_model
	collection := client.Database("GODB").Collection("review")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"movie_name": review.Movie_name})
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err = cursor.All(ctx, &filtered); ErrHandler.Log(err) {
		return nil, err
	}
	return filtered, nil
}

func (s *service) ShowReviewsBy_Email_and_MovieName(review *Models.Review_model) ([]Models.Review_model, error) {
	var filtered []Models.Review_model
	collection := client.Database("GODB").Collection("review")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"email": review.Email, "movie_name": review.Movie_name})
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err := cursor.All(ctx, &filtered); ErrHandler.Log(err) {
		return nil, err
	}
	return filtered, nil
}

func (s *service) Count_Reviews_and_Watchlist_ByEmail(review *Models.Account) (map[string]int64, error) {
	collection_review := client.Database("GODB").Collection("review")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor_review, err := collection_review.CountDocuments(ctx, bson.M{"email": review.Email})
	if ErrHandler.Log(err) {
		return nil, err
	}
	collection_watchlist := client.Database("GODB").Collection("watchlist")
	cursor_watchlist, err := collection_watchlist.CountDocuments(ctx, bson.M{"email": review.Email})
	if ErrHandler.Log(err) {
		return nil, err
	}
	result := map[string]int64{
		"review":    cursor_review,
		"watchlist": cursor_watchlist,
	}
	return result, nil
}
