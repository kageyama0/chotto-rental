package review_handler

import (
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
)

func (h *ReviewHandler) updateUserTrustScore(userID uuid.UUID) error {
	var averageScore float64
	err := h.db.Model(&model.Review{}).
		Where("reviewed_user_id = ?", userID).
		Select("COALESCE(AVG(score), 1.0)").
		Scan(&averageScore).Error
	if err != nil {
		return err
	}

	return h.db.Model(&model.User{}).
		Where("id = ?", userID).
		Update("trust_score", averageScore).Error
}
