package case_handler

import (
	"time"

	"gorm.io/gorm"
)

type CaseHandler struct {
	db *gorm.DB
}

type CreateCaseRequest struct {
	Title          string    `json:"title" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	Reward         int       `json:"reward" binding:"required,min=0"`
	Location       string    `json:"location" binding:"required"`
	ScheduledDate  time.Time   `json:"scheduledDate" binding:"required"`
	DurationMinutes int      `json:"durationMinutes" binding:"required,min=1"`
}
