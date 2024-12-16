package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
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
// @Success 201 {object} util.Response{data=AuthResponse} "登録成功"
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

	// メールアドレスの重複チェック
	var existingUser model.User
	if result := h.db.Where("email = ?", req.Email).First(&existingUser); result.Error == nil {
		util.CreateResponse(c, http.StatusConflict, e.EMAIL_ALREADY_EXISTS, nil)
		return
	}

	// パスワードハッシュ化
	hashedPassword, err := h.authService.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "パスワードの処理に失敗しました"})
		return
	}

	user := model.User{
		Email:        req.Email,
		PasswordHash: hashedPassword,
		DisplayName:  req.DisplayName,
		TrustScore:   1.0,
	}

	if result := h.db.Create(&user); result.Error != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	// JWTトークン生成
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
			"displayName": user.DisplayName,
			"trustScore":  user.TrustScore,
		},
	}

	util.CreateResponse(c, http.StatusCreated, e.CREATED, response)
}

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
