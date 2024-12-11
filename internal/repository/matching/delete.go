package matching_repository

import "github.com/google/uuid"

// -- DeleteByRequesterIDAndHelperID: あるユーザーが、リクエスター or ヘルパーとして関わっているマッチングを削除する
func (r *MatchingRepository) DeleteByRequesterIDAndHelperID(userID *uuid.UUID) error {
	err := r.db.Where("requester_id = ? AND helper_id = ?", userID, userID).Delete(&Matching{}).Error
	if err != nil {
		return err
	}

	return nil
}
