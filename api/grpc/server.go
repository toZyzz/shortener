package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"net"
	"shortener/domain/action/url"
	grpc_v1 "shortener/grpc.v1"
	"shortener/infrastructure/config"
)

type UrlShortenerServer struct {
	grpc_v1.UnimplementedUrlShortenerServer
}

func NewUrlShortenerServer() *UrlShortenerServer {
	return &UrlShortenerServer{}
}

func (s *UrlShortenerServer) ShortenUrl(ctx context.Context, req *grpc_v1.ShortenUrlRequest) (*grpc_v1.ShortenUrlResponse, error) {
	action := url.NewShortAction()
	shortUrl, err := action.GenerateAndSaveShortURL(req.FullUrl)
	if err != nil {
		return nil, err
	}
	return &grpc_v1.ShortenUrlResponse{ShortUrl: shortUrl}, nil
}

func (s *UrlShortenerServer) UnshortenUrl(ctx context.Context, req *grpc_v1.UnshortenUrlRequest) (*grpc_v1.UnshortenUrlResponse, error) {
	urlEntity, err := config.UrlRepository.GetByShortUrl(req.ShortUrl)

	if err != nil {
		return nil, err
	}

	return &grpc_v1.UnshortenUrlResponse{FullUrl: urlEntity.Full}, nil
}

func StartGRPCServer(port string) error {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	grpc_v1.RegisterUrlShortenerServer(grpcServer, NewUrlShortenerServer())

	log.Printf("gRPC server is running on port %s", port)
	return grpcServer.Serve(lis)
}
