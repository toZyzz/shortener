package url

import (
	"github.com/gin-gonic/gin"
	"shortener/infrastructure/config"
)

func Unshort(c *gin.Context) {
	shortUrl := c.Param("short_url")
	url, err := config.UrlRepository.GetByShortUrl(shortUrl)

	if err != nil {
		c.JSON(200, gin.H{
			"full": nil,
		})
		return
	}

	c.JSON(200, gin.H{
		"full": url.Full,
	})
}
