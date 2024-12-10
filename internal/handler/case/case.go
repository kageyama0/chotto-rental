package case_handler

import (
	"gorm.io/gorm"
)

type CaseHandler struct {
	db *gorm.DB
}

func NewCaseHandler(db *gorm.DB) *CaseHandler {
	return &CaseHandler{db: db}
}
