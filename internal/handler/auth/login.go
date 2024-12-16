package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
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

	var user model.User
	if result := h.db.Where("email = ?", req.Email).First(&user); result.Error != nil {
		util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_EMAIL_OR_PASSWORD, nil)
		return
	}

	if !h.authService.CheckPassword(req.Password, user.PasswordHash) {
		util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_EMAIL_OR_PASSWORD, nil)
		return
	}

	token, err := h.authService.GenerateToken(user.ID.String())
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	response := gin.H{
		"token": token,
		"user": gin.H{
			"id":           user.ID,
			"email":        user.Email,
			"display_name": user.DisplayName,
			"trust_score":  user.TrustScore,
		},
	}

	util.CreateResponse(c, http.StatusOK, e.OK, response)
}
