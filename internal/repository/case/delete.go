package case_repository

import (
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
)

// - DeleteByID
// idで案件を削除
func (r *CaseRepository) DeleteByID(id uuid.UUID) error {
	err := r.db.Delete(&model.Case{}, "id = ?", id).Error
	return err
}
