package review_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	_ "github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary レビュー一覧取得
// @Description ユーザーのレビュー一覧を取得します。user_idクエリパラメータを指定すると、特定のユーザーに対するレビューのみを取得します。
// @Tags レビュー
// @Accept json
// @Produce json
// @Param user_id query string false "ユーザーID（指定しない場合は全レビューを取得）"
// @Success 200 {array} model.Review "レビュー一覧"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /reviews [get]
func (h *ReviewHandler) List(c *gin.Context) {
	userID := c.Query("user_id")
	var reviews []model.Review

	query := h.db.Preload("Reviewer").Preload("ReviewedUser").Preload("Matching")
	if userID != "" {
		query = query.Where("reviewed_user_id = ?", userID)
	}

	if err := query.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "レビューの取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
