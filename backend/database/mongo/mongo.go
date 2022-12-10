package MongoDB

import (
	config "backend/config"
	ErrHandler "backend/helper_handlers/error"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

// Mongoclient open connection to mongo
func Mongoclient() error {
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Mongo_Host))
	if ErrHandler.Log(err) {
		return err
	}
	if err := client.Ping(context.TODO(), nil); ErrHandler.Log(err) {
		return err
	}
	return nil
}
