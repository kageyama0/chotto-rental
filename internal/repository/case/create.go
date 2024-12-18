package case_repository

// --Create: 案件を作成
func (r *CaseRepository) Create(c *Case) error {
	err := r.db.Create(c).Error
	return err
}
