package main

import (
	config "backend/config"
	MongoDB "backend/database/mongo"
	RedisDB "backend/database/redis"
	Server "backend/server"
	"log"
)

func main() {
	if err := MongoDB.Mongoclient(); err != nil {
		panic(err)
	}
	log.Println("Mongodb connected", config.Mongo_Host)
	if err := RedisDB.Redisclient(); err != nil {
		panic(err)
	}
	log.Println("Redis connected", config.Redis_Host)
	log.Println("API LISTEN ON", config.Server_Port)
	if err := Server.Start(config.Server_Port); err != nil {
		panic(err)
	}
}
