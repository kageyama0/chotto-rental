package review_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
)


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
