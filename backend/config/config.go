package config

import "os"

var (
	Server_Port    = os.Getenv("Server_Port")
	Redis_Host     = os.Getenv("Redis_Host") // Ex. Docker: redis:6379 | For local unit-test: localhost:6379
	Mongo_Host     = os.Getenv("Mongo_Host") // Ex. Docker: mongodb://mongodb:27017 | For local unit-test: mongodb://localhost:27017
	Access_Secret  = os.Getenv("Access_Secret")
	Refresh_Secret = os.Getenv("Refresh_Secret")
)
