package RedisDB

import (
	"backend/config"
	ErrHandler "backend/helper_handlers/error"

	"github.com/go-redis/redis/v7"
)

var client *redis.Client

// Redisclient open connection to redis
func Redisclient() error {
	client = redis.NewClient(&redis.Options{
		Addr: config.Redis_Host,
	})
	if _, err := client.Ping().Result(); ErrHandler.Log(err) {
		return err
	}
	return nil
}
