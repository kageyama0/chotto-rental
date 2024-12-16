package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

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
func (h *UserHandler) ListByUser(c *gin.Context) {
	_, userID, err := util.GetParams(c, []string{})
	if err != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	var reviews []model.Review
	if err := h.db.Preload("Reviewer").
		Where("reviewed_user_id = ?", userID).
		Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "レビューの取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
