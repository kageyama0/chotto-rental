package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
)

func (h *CaseHandler) List(c *gin.Context) {
	var cases []model.Case
	if err := h.db.Preload("User").Find(&cases).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "案件の取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, cases)
}

func (h *CaseHandler) Get(c *gin.Context) {
	id := c.Param("id")
	var caseData model.Case

	if err := h.db.Preload("User").First(&caseData, "id = ?", id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "案件が見つかりません"})
		return
	}

	c.JSON(http.StatusOK, caseData)
}
