package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	auth_service "github.com/kageyama0/chotto-rental/internal/service/auth"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

func AuthMiddleware(authService auth_service.AuthService) gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, errCode := authService.GetUserIDBySessionID(c)
		if errCode != e.OK {
			util.CreateResponse(c, http.StatusUnauthorized, errCode, nil)
			return
		}
		c.Set("userID", *userID)
		c.Next()
	}
}
