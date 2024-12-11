package matching_handler

import (
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)


type Matching = model.Matching
type Application = model.Application
type Case = model.Case
type MatchingHandler struct {
	db *gorm.DB
}

func NewMatchingHandler(db *gorm.DB) *MatchingHandler {
	return &MatchingHandler{db: db}
}
