package matching_repository

import (
	"github.com/kageyama0/chotto-rental/internal/model"

	"gorm.io/gorm"
)

type Matching = model.Matching
type MatchingRepository struct {
	db *gorm.DB
}

func NewMatchingRepository(db *gorm.DB) *MatchingRepository {
	return &MatchingRepository{db: db}
}
