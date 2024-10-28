package url

import (
	"math/rand"
	urlEntity "shortener/domain/entity/url"
	"shortener/infrastructure/config"
	"strings"
	"time"
)

type ShortAction struct{}

func NewShortAction() *ShortAction {
	return &ShortAction{}
}

func (s *ShortAction) GenerateAndSaveShortURL(fullUrl string) (string, error) {
	url, _ := config.UrlRepository.GetByFullUrl(fullUrl)

	if url != nil {
		return url.Short, nil
	}

	url = &urlEntity.Url{Short: generateShortURL(), Full: fullUrl}
	err := config.UrlRepository.Save(url)

	return url.Short, err
}

func generateShortURL() string {
	rand.Seed(time.Now().UnixNano())
	charset := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	var sb strings.Builder

	for i := 0; i < 10; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}

	return sb.String()
}
