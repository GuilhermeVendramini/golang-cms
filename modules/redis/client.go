package redis

import (
	"github.com/go-redis/redis"
)

// Client Redis
var Client *redis.Client

// Nil redis
var Nil error

func init() {
	Client = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	Nil = redis.Nil
}
