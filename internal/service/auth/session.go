package auth_service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

// cookieにセットされたセッションIDからuserIDを取得する
func (s *AuthService) GetUserIDBySessionID(c *gin.Context) (userID *uuid.UUID, errCode int) {
	// クッキーからセッションIDを取得
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		c.Abort()
		return nil, e.NOT_FOUND_SESSION_ID
	}

	// セッションIDのパース
	parsedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		c.Abort()
		return nil, e.INVALID_SESSION_ID
	}

	// セッションの検証
	session, err := s.sessionRepository.FindByID(parsedSessionID)
	if err != nil {
		c.Abort()
		return nil, e.NOT_FOUND_SESSION
	}

	// セッションの有効性チェック
	if !session.IsValid || time.Now().After(session.ExpiresAt) {
		c.Abort()
		return nil, e.INVALID_SESSION
	}

	userID = &session.UserID

	return userID, e.OK
}
