package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	_ "github.com/kageyama0/chotto-rental/pkg/util"
)

type UpdateUserRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
}


// @Summary ユーザー情報更新
// @Description ログインユーザーの表示名を更新します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param request body UpdateUserRequest true "更新するユーザー情報"
// @Success 200 {object} map[string]interface{} "更新後のユーザー情報"
// @Failure 400 {object} util.Response "リクエストが不正です"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 404 {object} util.Response "ユーザーが見つかりません"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /users/me [put]
func (h *UserHandler) Update(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid, _ := uuid.Parse(userID.(string))

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := h.db.First(&user, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	user.DisplayName = req.DisplayName
	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザー情報の更新に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"email":        user.Email,
		"display_name": user.DisplayName,
		"trust_score":  user.TrustScore,
	})
}
