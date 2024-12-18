package auth_service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// Signup ユーザー登録処理
func (s *AuthService) Signup(c *gin.Context, email, password string, deviceInfo DeviceInfo) (session *Session, statusCode int, errCode int) {
	// メールアドレスが既に登録されていないか確認
	_, errCode = s.userRepository.FindByEmail(email)
	if errCode == e.SERVER_ERROR {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}
	if errCode != e.NOT_FOUND_USER {
		util.CreateResponse(c, http.StatusBadRequest, e.ALREADY_REGISTERED_EMAIL, nil)
		return
	}

	// パスワードのハッシュ化
	hashedPassword, err := s.hashPassword(password)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	// トランザクション開始
	tx := s.db.Begin()
	if tx.Error != nil {
		return nil, http.StatusInternalServerError, e.SERVER_ERROR
	}
	defer tx.Rollback()

	// ユーザーの作成
	user := &User{
		ID:           uuid.New(),
		Email:        email,
		PasswordHash: string(hashedPassword),
	}
	if err := s.userRepository.CreateWithTransaction(tx, user); err != nil {
		return nil, http.StatusInternalServerError, e.SERVER_ERROR
	}

	// セッションの作成
	session = &Session{
		UserID:         user.ID,
		DeviceInfo:     deviceInfo,
		LastAccessedAt: time.Now(),
		// TODO: condigで指定するようにする。一旦30日
		ExpiresAt: time.Now().Add(30 * 24 * time.Hour),
		IsValid:   true,
	}
	if err := s.sessionRepository.CreateWithTransaction(tx, session); err != nil {
		return nil, http.StatusInternalServerError, e.SERVER_ERROR
	}

	// トランザクションのコミット
	if err := tx.Commit().Error; err != nil {
		return nil, http.StatusInternalServerError, e.SERVER_ERROR
	}

	return session, http.StatusOK, e.OK
}
