package case_repository

import "github.com/google/uuid"

// -- UpdateStatus: 案件のステータスを更新する
func (r *CaseRepository) UpdateStatus(caseID uuid.UUID, status string) error {
	var c Case
	err := r.db.Model(&c).Where("id = ?", caseID).Update("status", status).Error
	if err != nil {
		return err
	}

	return nil
}
