package handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type CaseHandler struct {
	db *gorm.DB
}

func NewCaseHandler(db *gorm.DB) *CaseHandler {
	return &CaseHandler{db: db}
}

type CreateCaseRequest struct {
	Title          string    `json:"title" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	Reward         int       `json:"reward" binding:"required,min=0"`
	Location       string    `json:"location" binding:"required"`
	ScheduledDate  time.Time   `json:"scheduled_date" binding:"required"`
	DurationMinutes int      `json:"duration_minutes" binding:"required,min=1"`
}

func (h *CaseHandler) Create(c *gin.Context) {
	var req CreateCaseRequest
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

	caseData := model.Case{
		UserID:         uid,
		Title:          req.Title,
		Description:    req.Description,
		Reward:         req.Reward,
		Location:       req.Location,
		ScheduledDate:  req.ScheduledDate,
		DurationMinutes: req.DurationMinutes,
		Status:         "open",
	}

	if result := h.db.Create(&caseData); result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "案件の作成に失敗しました"})
		return
	}

	c.JSON(http.StatusCreated, caseData)
}

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
