package matching_repository

import (
	"github.com/google/uuid"
)

// -- FindByID: IDを使用してマッチングを取得する
func (r *MatchingRepository) FindByID(id uuid.UUID) (Matching, error) {
	var matching Matching

	err := r.db.First(&matching, "id = ?", id).Error

	if err != nil {
		return matching, err
	}

	return matching, nil
}
