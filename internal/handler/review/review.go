package review_handler

import (
	"gorm.io/gorm"
)

type ReviewHandler struct {
	db *gorm.DB
}

func NewReviewHandler(db *gorm.DB) *ReviewHandler {
	return &ReviewHandler{db: db}
}
