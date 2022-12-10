package MongoDB

import (
	ErrHandler "backend/helper_handlers/error"
	Models "backend/models"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (s *service) AddWatchlist(watchlist *Models.Watchlist) error {
	var dbwatchlist *Models.Watchlist
	collection := client.Database("GODB").Collection("watchlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if ifexist := collection.FindOne(ctx, bson.M{"movie_name": watchlist.Movie_name}).Decode(&dbwatchlist); ifexist == nil {
		ErrHandler.Log(fmt.Errorf("watchlist movie already exist"))
		return fmt.Errorf("watchlist movie already exist")
	}
	watchlist, err := InitWatchlistBSON(watchlist)
	if ErrHandler.Log(err) {
		return err
	}
	if _, err := collection.InsertOne(ctx, watchlist); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) DeleteWatchlist(watchlist *Models.DeleteWatchlistMany) error {
	collection := client.Database("GODB").Collection("watchlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	delete := bson.M{"email": watchlist.Email, "movie_name": bson.M{"$in": watchlist.Movie_name}}
	if _, err := collection.DeleteMany(ctx, delete); ErrHandler.Log(err) {
		return err
	}
	return nil
}

func (s *service) ShowWatchlistByEmail(watchlist *Models.Watchlist) ([]Models.Watchlist, error) {
	var filtered []Models.Watchlist
	collection := client.Database("GODB").Collection("watchlist")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	cursor, err := collection.Find(ctx, bson.M{"email": watchlist.Email})
	if ErrHandler.Log(err) {
		return nil, err
	}
	if err = cursor.All(ctx, &filtered); ErrHandler.Log(err) {
		return nil, err
	}
	return filtered, nil
}
