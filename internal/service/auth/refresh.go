package auth_service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

func (s *AuthService) Refresh(c *gin.Context, sessionID uuid.UUID, deviceInfo DeviceInfo) (statusCode int, errCode int) {
	session, errCode := s.sessionRepository.FindByID(sessionID)
	if errCode != e.OK {
		return http.StatusUnauthorized, errCode
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

	_, errCode = s.userRepository.FindByID(session.UserID)
	if errCode != e.OK {
		return http.StatusInternalServerError, errCode
	}

	return http.StatusOK, e.OK
}
