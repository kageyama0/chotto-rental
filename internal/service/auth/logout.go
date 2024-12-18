package auth_service

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

// Logout; ログアウト処理
func (s *AuthService) Logout(c *gin.Context, sessionID uuid.UUID) (statusCode int, errCode int) {
    // セッションの無効化
    if err := s.sessionRepository.Invalidate(sessionID); err != nil {
        return http.StatusInternalServerError, e.SERVER_ERROR
    }

    return http.StatusOK, e.OK
}
