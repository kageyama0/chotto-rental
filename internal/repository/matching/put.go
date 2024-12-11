package matching_repository

// -- Update: マッチングを更新する
func (r *MatchingRepository) Update(matching *Matching) error {
	err := r.db.Save(matching).Error
	if err != nil {
		return err
	}

	return nil
}
