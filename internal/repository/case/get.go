package case_repository

import (
	"gorm.io/gorm"

	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
)

type Case = model.Case
type CaseRepository struct {
	db *gorm.DB
}

func NewCaseRepository(db *gorm.DB) *CaseRepository {
	return &CaseRepository{db: db}
}

// idで案件を取得
func (r *CaseRepository) FindByID(id uuid.UUID) (*Case, error) {
	var c Case

	if err := r.db.First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &c, nil
}
