package case_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
)

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
