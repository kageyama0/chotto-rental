package case_handler

import (
	"time"

	"gorm.io/gorm"
)

type CaseHandler struct {
	db *gorm.DB
}

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
	Address        string    `json:"address"`
}
