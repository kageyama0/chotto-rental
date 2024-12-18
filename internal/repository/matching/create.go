package matching_repository

// -- Create: マッチングを作成する
func (r *MatchingRepository) Create(matching *Matching) error {
	err := r.db.Create(matching).Error
	if err != nil {
		return err
	}

	return nil
}
