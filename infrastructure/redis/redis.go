package redis

import (
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var Rdb *redis.Client

func ConnectToRedis() {
	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic("Couldn`t parse REDIS_DB to int: " + err.Error())
	}

	Rdb = redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_URL"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
}
