package auth_service

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// cookieにセットされたセッションIDからuserIDを取得する
func (s *AuthService) GetUserIDBySessionID(c *gin.Context) (session *Session, errCode int) {
		// クッキーからセッションIDを取得
		sessionID, err := c.Cookie("session_id")
		if err != nil {
				c.Abort()
				util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_SESSION_ID, nil)
				return
		}

		// セッションIDのパース
		parsedSessionID, err := uuid.Parse(sessionID)
		if err != nil {
				c.Abort()
				util.CreateResponse(c, http.StatusBadRequest, e.INVALID_SESSION_ID, nil)
				return
		}

		// セッションの検証
		session, err = s.sessionRepository.FindByID(parsedSessionID)
		if err != nil {
				c.Abort()
				util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_SESSION_ID, nil)
				return
		}

		// セッションの有効性チェック
		if !session.IsValid || time.Now().After(session.ExpiresAt) {
			  c.Abort()
				util.CreateResponse(c, http.StatusUnauthorized, e.SESSION_EXPIRED, nil)
				return
		}

		return session, e.OK
}
