package review_repository

import (
	"github.com/google/uuid"
)

// DeleteByReviewerIDAndReviewedID: あるユーザーが書いた or 書かれたレビューを削除する
func (r *ReviewRepository) DeleteByReviewerIDAndReviewedID(userID *uuid.UUID) error {
	err := r.db.Where("reviewer_id = ? AND reviewed_id = ?", userID, userID).Delete(&Review{}).Error
	if err != nil {
		return err
	}

	return nil
}
