package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// 登録ハンドラー
// @Summary ユーザー登録
// @Description 新規ユーザーを登録し、認証トークンを発行します
// @Tags 認証
// @Accept json
// @Produce json
// @Param request body SignupRequest true "登録情報"
// @Success 201 {object} util.Response "登録成功"
// @Failure 400 {object} util.Response "無効なパラメータ"
// @Failure 409 {object} util.Response "メールアドレスが既に使用されています"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /auth/signup [post]
func (h *AuthHandler) Signup(c *gin.Context) {
	var req SignupRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.JSON_PARSE_ERROR, nil)
		return
	}

	deviceInfo := DeviceInfo{
		UserAgent: c.GetHeader("User-Agent"),
		IP:        c.ClientIP(),
	}

	session, statusCode, errCode := h.authService.Signup(c, req.Email, req.Password, deviceInfo)
	if errCode != e.OK {
		util.CreateResponse(c, statusCode, errCode, nil)
		return
	}
	if session == nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	// セッションIDをクッキーにセット
	c.SetCookie(
		"session_id",
		session.ID.String(),
		// TODO: configで指定するようにする。一旦30日
		30*24*60*60,
		"/",
		"",
		true,
		true,
	)

	util.CreateResponse(c, http.StatusCreated, e.CREATED, nil)
}
