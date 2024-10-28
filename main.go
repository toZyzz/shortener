package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"shortener/api/grpc"
	"shortener/api/http/url"
	"shortener/infrastructure/config"
)

func init() {
	config.LoadEnvVariables()
	config.LoadRepository()
}

func main() {
	go func() {
		if err := grpc.StartGRPCServer(":50051"); err != nil {
			log.Fatalf("Failed to start gRPC server: %v", err)
		}
	}()

	r := gin.Default()
	r.POST("/short", url.Short)
	r.GET("/full/:short_url", url.Unshort)
	r.Run()
}
