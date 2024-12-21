package case_repository

import "github.com/google/uuid"

// --Create: 案件を作成
func (r *CaseRepository) Create(c *Case) (uuid.UUID, error) {
	err := r.db.Create(c).Error
	return c.ID, err
}
