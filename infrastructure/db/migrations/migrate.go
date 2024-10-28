package main

import (
	"shortener/domain/entity/url"
	"shortener/infrastructure/config"
	"shortener/infrastructure/db"
)

func init() {
	config.LoadEnvVariables()
	db.ConnectToDB()
}

func main() {
	db.DB.AutoMigrate(&url.Url{})
}
