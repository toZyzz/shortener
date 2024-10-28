package repository

import (
	"shortener/domain/entity/url"
	"shortener/infrastructure/db"
)

type DatabaseURLRepository struct{}

func NewDatabaseURLRepository() *DatabaseURLRepository {
	return &DatabaseURLRepository{}
}

func (r *DatabaseURLRepository) Save(url *url.Url) error {
	result := db.DB.Create(url)

	return result.Error
}

func (r *DatabaseURLRepository) GetByFullUrl(fullUrl string) (*url.Url, error) {
	var url url.Url
	result := db.DB.Where(`"full" = ?`, fullUrl).First(&url)

	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}

func (r *DatabaseURLRepository) GetByShortUrl(shortUrl string) (*url.Url, error) {
	var url url.Url
	result := db.DB.Where("short = ?", shortUrl).First(&url)

	if result.Error != nil {
		return nil, result.Error
	}

	return &url, nil
}
