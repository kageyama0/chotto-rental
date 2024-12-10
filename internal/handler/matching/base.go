package matching_handler

import (
	"gorm.io/gorm"
)

type MatchingHandler struct {
	db *gorm.DB
}

func NewMatchingHandler(db *gorm.DB) *MatchingHandler {
	return &MatchingHandler{db: db}
}
