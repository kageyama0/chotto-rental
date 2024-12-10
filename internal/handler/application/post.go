package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
)


type CreateApplicationRequest struct {
	CaseID  string `json:"case_id" binding:"required"`
	Message string `json:"message"`
}

func (h *ApplicationHandler) Create(c *gin.Context) {
	var req CreateApplicationRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	uid, err := uuid.Parse(userID.(string))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーID"})
		return
	}

	caseID, err := uuid.Parse(req.CaseID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な案件ID"})
		return
	}

	var caseData model.Case
	if err := h.db.First(&caseData, "id = ?", caseID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "案件が見つかりません"})
		return
	}

	if caseData.Status != "open" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "この案件は応募を受け付けていません"})
		return
	}

	var existingApplication model.Application
	err = h.db.Where("case_id = ? AND applicant_id = ?", caseID, uid).First(&existingApplication).Error
	if err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "既にこの案件に応募しています"})
		return
	}

	application := model.Application{
		CaseID:      caseID,
		ApplicantID: uid,
		Message:     req.Message,
		Status:      "pending",
	}

	if result := h.db.Create(&application); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "応募の作成に失敗しました"})
		return
	}

	c.JSON(http.StatusCreated, application)
}
