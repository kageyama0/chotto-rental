package matching_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
)

func (h *MatchingHandler) ConfirmArrival(c *gin.Context) {
	matchingID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なマッチングID"})
		return
	}

	var matching model.Matching
	if err := h.db.First(&matching, "id = ?", matchingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "マッチングが見つかりません"})
		return
	}

	userID, _ := c.Get("userID")
	uid := userID.(string)

	if matching.Status != "active" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "このマッチングは確認できません"})
		return
	}

	if time.Now().After(matching.ArrivalConfirmationDeadline) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "確認期限が過ぎています"})
		return
	}

	switch uid {
	case matching.RequesterID.String():
		matching.ArrivalConfirmedByRequester = true
	case matching.HelperID.String():
		matching.ArrivalConfirmedByHelper = true
	default:
		c.JSON(http.StatusForbidden, gin.H{"error": "この操作を行う権限がありません"})
		return
	}

	if matching.ArrivalConfirmedByRequester && matching.ArrivalConfirmedByHelper {
		matching.Status = "completed"
	}

	if err := h.db.Save(&matching).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "到着確認の更新に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, matching)
}
