package case_repository

import (
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type Case = model.Case
type CaseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) *CaseRepository {
	return &CaseRepository{db: db}
}
