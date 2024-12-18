package auth_service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

func (s *AuthService) Refresh(c *gin.Context, sessionID uuid.UUID, deviceInfo DeviceInfo) (statusCode int, errCode int) {
	session, err := s.sessionRepository.FindByID(sessionID)
	if err != nil {
		return http.StatusUnauthorized, e.NOT_FOUND_SESSION
	}

	// セッションの有効期限チェック
	if time.Now().After(session.ExpiresAt) {
		s.sessionRepository.Invalidate(sessionID)
		return http.StatusUnauthorized, e.SESSION_EXPIRED
	}

	// デバイス情報の検証
	if !validateDeviceInfo(session.DeviceInfo, deviceInfo) {
		s.sessionRepository.Invalidate(sessionID)
		return http.StatusUnauthorized, e.INVALID_DEVICE
	}

	user, err := s.userRepository.FindByID(session.UserID)
	if err != nil {
		return http.StatusInternalServerError, e.SERVER_ERROR
	}
	if user == nil {
		return http.StatusUnauthorized, e.NOT_FOUND_USER
	}

	return http.StatusOK, e.OK
}
