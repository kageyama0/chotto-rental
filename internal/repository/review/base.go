package review_repository

import (
	"github.com/kageyama0/chotto-rental/internal/model"

	"gorm.io/gorm"
)

type Review = model.Review
type ReviewRepository struct {
	db *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{db: db}
}
