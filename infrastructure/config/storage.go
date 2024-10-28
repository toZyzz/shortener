package config

import (
	"os"
	repositoryContract "shortener/domain/contracts/repository"
	"shortener/infrastructure/db"
	"shortener/infrastructure/redis"
	"shortener/infrastructure/repository"
)

var UrlRepository repositoryContract.UrlRepository

func LoadRepository() {
	mode := os.Getenv("MODE")

	if mode == "in_memory" {
		redis.ConnectToRedis()
		UrlRepository = repository.NewInMemoryURLRepository()
	}

	if mode == "db" {
		db.ConnectToDB()
		UrlRepository = repository.NewDatabaseURLRepository()
	}

}
