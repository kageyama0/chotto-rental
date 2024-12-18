package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// ログインハンドラー
// @Summary ログイン
// @Description メールアドレスとパスワードで認証し、トークンを発行します
// @Tags 認証
// @Accept json
// @Produce json
// @Param request body LoginRequest true "ログイン情報"
// @Success 200 {object} util.Response{data=AuthResponse} "ログイン成功"
// @Failure 400 {object} util.Response "無効なパラメータ"
// @Failure 401 {object} util.Response "メールアドレスまたはパスワードが間違っています"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /auth/login [post]
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.JSON_PARSE_ERROR, nil)
		return
	}

	deviceInfo := DeviceInfo{
			UserAgent: c.GetHeader("User-Agent"),
			IP:        c.ClientIP(),
	}

	session, errCode := h.authService.Login(c, req.Email, req.Password, deviceInfo)
	if errCode != e.OK {
			util.CreateResponse(c, http.StatusUnauthorized, errCode, nil)
			return
	}

	// セッションIDをクッキーにセット
	c.SetCookie(
			"session_id",
			session.ID.String(),
			30*24*60*60, // 30日
			"/",
			"",
			true,
			true,
	)

	util.CreateResponse(c, http.StatusOK, e.OK, nil)
}
