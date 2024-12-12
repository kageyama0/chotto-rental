package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	_ "github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary ユーザー情報取得
// @Description ユーザーIDに基づいてユーザーの基本情報を取得します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} map[string]interface{} "ユーザー情報"
// @Failure 400 {object} util.Response "無効なユーザーID"
// @Failure 404 {object} util.Response "ユーザーが見つかりません"
// @Router /users/{id} [get]
func (h *UserHandler) Get(c *gin.Context) {
	userID := c.Param("id")
	uid, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーID"})
		return
	}

	var user model.User
	if err := h.db.First(&user, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"email":        user.Email,
		"display_name": user.DisplayName,
		"trust_score":  user.TrustScore,
	})
}

// @Summary ユーザーのレビュー一覧取得
// @Description 指定されたユーザーが受け取ったレビューの一覧を取得します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {array} model.Review "レビュー一覧"
// @Failure 400 {object} util.Response "無効なユーザーID"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /users/{id}/reviews [get]
func (h *UserHandler) GetReviews(c *gin.Context) {
	userID := c.Param("id")
	uid, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーID"})
		return
	}

	var reviews []model.Review
	if err := h.db.Preload("Reviewer").
		Where("reviewed_user_id = ?", uid).
		Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "レビューの取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
