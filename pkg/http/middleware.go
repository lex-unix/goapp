package http

import (
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Authentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		sessionId := session.Get("id")
		if sessionId == nil {
			c.JSON(http.StatusUnauthorized, gin.H{
				"message": "unathorized",
			})
			c.Abort()
		}
	}
}
