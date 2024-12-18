package case_repository

import (
	"github.com/google/uuid"
)

// -- FindByID: idで案件を取得
func (r *CaseRepository) FindByID(id uuid.UUID) (*Case, error) {
	var c Case

	err := r.db.First(&c, "id = ?", id).Error
	if err != nil {
		return nil, err
	}

	return &c, nil
}

// -- FindAll: すべての案件を取得
func (r *CaseRepository) FindAll() ([]Case, error) {
	var cases []Case

	err := r.db.Preload("User").Find(&cases).Error
	if err != nil {
		return nil, err
	}

	return cases, nil
}

// -- FindByIdWithUser: idで案件を取得し、ユーザー情報も取得
func (r *CaseRepository) FindByIDWithUser(id uuid.UUID) (*Case, error) {
	var c Case

	if err := r.db.Preload("User").First(&c, "id = ?", id).Error; err != nil {
		return nil, err
	}

	return &c, nil
}
