package middlewares

import (
	"ecommerce/models"
	"github.com/gin-gonic/gin"

	"net/http"
)

func EnforceAuthenticatedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, exists := c.Get("currentUser")
		if exists && user.(models.User).ID != 0 {
			return
		} else {
			err, _ := c.Get("authErr")
			_ = c.AbortWithError(http.StatusUnauthorized, err.(error))
			return
		}
	}
}
