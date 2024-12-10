package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
)

type UpdateApplicationStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=accepted rejected"`
}

func (h *ApplicationHandler) UpdateStatus(c *gin.Context) {
	id := c.Param("id")
	var req UpdateApplicationStatusRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var application model.Application
	if err := h.db.Preload("Case").First(&application, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "応募が見つかりません"})
		return
	}

	userID, _ := c.Get("userID")
	if application.Case.UserID.String() != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "この操作を行う権限がありません"})
		return
	}

	application.Status = req.Status
	if err := h.db.Save(&application).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "応募状態の更新に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, application)
}
