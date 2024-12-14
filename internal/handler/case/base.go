package case_handler

import (
	"gorm.io/gorm"
)

func NewCaseHandler(db *gorm.DB) *CaseHandler {
	return &CaseHandler{db: db}
}
