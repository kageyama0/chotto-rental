package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/pkg/service"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

func AuthMiddleware(authService *service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			util.CreateResponse(c, http.StatusUnauthorized, e.AUTH_REQUIRED, nil)
			c.Abort()
			return
		}

		bearerToken := strings.Split(authHeader, " ")
		if len(bearerToken) != 2 {
			util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_TOKEN_FORMAT, nil)
			c.Abort()
			return
		}

		claims, err := authService.ValidateToken(bearerToken[1])
		if err != nil {
			util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_TOKEN, nil)
			c.Abort()
			return
		}

		c.Set("userID", claims.UserID)
		c.Next()
	}
}
