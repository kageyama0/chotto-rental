package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
)

func (h *ApplicationHandler) List(c *gin.Context) {
	userID, _ := c.Get("userID")
	var applications []model.Application

	if err := h.db.Preload("Case").Preload("Applicant").
		Where("applicant_id = ?", userID).
		Find(&applications).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "応募の取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, applications)
}
