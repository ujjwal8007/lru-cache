package middleware

import (
	"encoding/base64"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func BasicAuthMiddleware(username, password string) gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := strings.SplitN(c.GetHeader("Authorization"), " ", 2)

		if len(auth) != 2 || auth[0] != "Basic" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization failed"})
			return
		}

		payload, _ := base64.StdEncoding.DecodeString(auth[1])
		pair := strings.SplitN(string(payload), ":", 2)

		if len(pair) != 2 || pair[0] != username || pair[1] != password {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "authorization failed"})
			return
		}

		c.Next()
	}
}
