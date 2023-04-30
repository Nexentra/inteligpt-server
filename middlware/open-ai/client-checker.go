package openai

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/nexentra/inteligpt/pkg/common/settings"
)

func KeyRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		key := c.GetHeader("Key")
		if key == "" {
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}

		c.Set("key", key)
		settings.SetOpenAiClient(key)
		c.Next()
	}
}
