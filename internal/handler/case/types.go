package case_handler

import (
	"time"

	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type CaseHandler struct {
	db *gorm.DB
}

type Case = model.Case

type CreateCaseRequest struct {
	Title          string    `json:"title" binding:"required,max=100"`
	Description    string    `json:"description" binding:"required,max=2000"`
	Category       string    `json:"category" binding:"required"`
	Reward         int       `json:"reward" binding:"required,min=500,max=100000"`
	RequiredPeople int       `json:"requiredPeople" binding:"required,min=1,max=10"`
	ScheduledDate  time.Time `json:"scheduledDate" binding:"required"`
	StartTime      string    `json:"startTime" binding:"required"`
	Duration       int       `json:"duration" binding:"required,min=15,max=360"`
	Prefecture     string    `json:"prefecture" binding:"required"`
	City           string    `json:"city" binding:"required"`
	Address        string    `json:"address" binding:"required"`
}

type CreateCaseResponse struct {
	ID             uuid.UUID `json:"id"`
	Title          string    `json:"title"`
	Description    string    `json:"description"`
	Category       string    `json:"category"`
	Reward         int       `json:"reward"`
	RequiredPeople int       `json:"requiredPeople"`
	ScheduledDate  time.Time `json:"scheduledDate"`
	StartTime      string    `json:"startTime"`
	Duration       int       `json:"duration"`
	Prefecture     string    `json:"prefecture"`
	City           string    `json:"city"`
	Address        string    `json:"address"`
	Status         string    `json:"status"`
}
