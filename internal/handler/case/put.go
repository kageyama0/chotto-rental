package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
)


func (h *CaseHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req CreateCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var caseData model.Case
	if err := h.db.First(&caseData, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "案件が見つかりません"})
		return
	}

	userID, _ := c.Get("userID")
	if caseData.UserID.String() != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "この操作を行う権限がありません"})
		return
	}

	caseData.Title = req.Title
	caseData.Description = req.Description
	caseData.Reward = req.Reward
	caseData.Location = req.Location
	caseData.ScheduledDate = req.ScheduledDate
	caseData.DurationMinutes = req.DurationMinutes

	if err := h.db.Save(&caseData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "案件の更新に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, caseData)
}
