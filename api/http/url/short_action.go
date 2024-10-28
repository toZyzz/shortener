package url

import (
	"github.com/gin-gonic/gin"
	actionUrl "shortener/domain/action/url"
)

func Short(c *gin.Context) {
	var body struct {
		Full string
	}

	c.Bind(&body)

	action := actionUrl.NewShortAction()
	shortUrl, err := action.GenerateAndSaveShortURL(body.Full)

	if err != nil {
		c.Status(400)
		return
	}

	c.JSON(201, gin.H{
		"short": shortUrl,
	})
}
