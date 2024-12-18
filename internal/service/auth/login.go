package auth_service

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

// Login: ログイン処理
func (s *AuthService) Login(c *gin.Context, email, password string, deviceInfo DeviceInfo) (session *Session, errCode int) {
	user, errCode := s.userRepository.FindByEmail(email)
	if errCode != e.OK {
		return nil, errCode
	}

	// パスワードが一致しているか確認
	if !s.checkPassword(password, user.PasswordHash) {
		return nil, e.INVALID_PASSWORD
	}

	// セッションの作成（有効期限は30日）
	session = &Session{
		UserID:         user.ID,
		DeviceInfo:     deviceInfo,
		LastAccessedAt: time.Now(),
		ExpiresAt:      time.Now().Add(30 * 24 * time.Hour),
		IsValid:        true,
	}

	// セッション情報の保存
	if err := s.sessionRepository.Create(session); err != nil {
		return nil, e.SERVER_ERROR
	}

	return session, e.OK
}
