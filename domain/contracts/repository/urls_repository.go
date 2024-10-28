package repository

import "shortener/domain/entity/url"

type UrlRepository interface {
	Save(post *url.Url) error
	GetByShortUrl(shortUrl string) (*url.Url, error)
	GetByFullUrl(fullUrl string) (*url.Url, error)
}
