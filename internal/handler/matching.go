package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type MatchingHandler struct {
	db *gorm.DB
}

func NewMatchingHandler(db *gorm.DB) *MatchingHandler {
	return &MatchingHandler{db: db}
}

type CreateMatchingRequest struct {
	ApplicationID   string `json:"application_id" binding:"required"`
	MeetingLocation string `json:"meeting_location" binding:"required"`
}

func (h *MatchingHandler) Create(c *gin.Context) {
	var req CreateMatchingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	applicationID, err := uuid.Parse(req.ApplicationID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効な応募ID"})
		return
	}

	var application model.Application
	if err := h.db.Preload("Case").First(&application, "id = ?", applicationID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "応募が見つかりません"})
		return
	}

	userID, _ := c.Get("userID")
	if application.Case.UserID.String() != userID.(string) {
		c.JSON(http.StatusForbidden, gin.H{"error": "この操作を行う権限がありません"})
		return
	}

	matching := model.Matching{
		CaseID:                      application.CaseID,
		RequesterID:                 application.Case.UserID,
		HelperID:                    application.ApplicantID,
		MeetingLocation:             req.MeetingLocation,
		ArrivalConfirmationDeadline: application.Case.ScheduledDate.Add(15 * time.Minute),
		Status:                      "active",
	}

	if err := h.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(&matching).Error; err != nil {
			return err
		}

		if err := tx.Model(&application).Update("status", "accepted").Error; err != nil {
			return err
		}

		if err := tx.Model(&model.Case{}).Where("id = ?", application.CaseID).Update("status", "matched").Error; err != nil {
			return err
		}

		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "マッチングの作成に失敗しました"})
		return
	}

	c.JSON(http.StatusCreated, matching)
}

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
