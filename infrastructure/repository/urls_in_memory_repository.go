package repository

import (
	"context"
	"errors"
	"shortener/domain/entity/url"
	redisInfra "shortener/infrastructure/redis"
)

type InMemoryUrlRepository struct{}

var ctx = context.Background()

func NewInMemoryURLRepository() *InMemoryUrlRepository {
	return &InMemoryUrlRepository{}
}

func (r *InMemoryUrlRepository) Save(url *url.Url) error {
	err := redisInfra.Rdb.Set(ctx, url.Full, url.Short, 0).Err()

	if err != nil {
		return err
	}

	return nil
}

func (r *InMemoryUrlRepository) GetByFullUrl(fullUrl string) (*url.Url, error) {
	shortUrl, err := redisInfra.Rdb.Get(ctx, fullUrl).Result()

	if err != nil {
		return nil, err
	}

	return &url.Url{Full: fullUrl, Short: shortUrl}, nil
}

func (r *InMemoryUrlRepository) GetByShortUrl(shortUrl string) (*url.Url, error) {
	iter := redisInfra.Rdb.Scan(ctx, 0, "", 0).Iterator()

	for iter.Next(ctx) {
		fullUrl := iter.Val()
		foundShortUrl, err := redisInfra.Rdb.Get(ctx, fullUrl).Result()
		if err == nil && foundShortUrl == shortUrl {
			return &url.Url{Full: fullUrl, Short: shortUrl}, nil
		}
	}

	if err := iter.Err(); err != nil {
		return nil, err
	}

	return nil, errors.New("Not found")
}
