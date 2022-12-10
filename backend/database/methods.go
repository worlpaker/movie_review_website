package db

import (
	MongoDB "backend/database/mongo"
	RedisDB "backend/database/redis"
)

type Handle struct {
	Redis RedisDB.IHandlers
	Mongo MongoDB.IHandlers
}

func NewAuth(m MongoDB.IHandlers, r RedisDB.IHandlers) *Handle {
	return &Handle{Redis: r, Mongo: m}
}

func NewProfile() *Handle {
	m := MongoDB.NewAuth()
	r := RedisDB.NewAuth()
	return NewAuth(m, r)
}

var Handler = NewProfile()
