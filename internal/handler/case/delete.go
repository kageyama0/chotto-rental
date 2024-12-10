package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
)

func (h *CaseHandler) Delete(c *gin.Context) {
	id := c.Param("id")
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

	if err := h.db.Delete(&caseData).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "案件の削除に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "案件を削除しました"})
}
