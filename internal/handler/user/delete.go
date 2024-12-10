package user_handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

func (h *UserHandler) Delete(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid, _ := uuid.Parse(userID.(string))

	if err := h.db.Transaction(func(tx *gorm.DB) error {
		// 関連データの削除
		if err := tx.Where("reviewer_id = ? OR reviewed_user_id = ?", uid, uid).Delete(&model.Review{}).Error; err != nil {
			return err
		}
		if err := tx.Where("requester_id = ? OR helper_id = ?", uid, uid).Delete(&model.Matching{}).Error; err != nil {
			return err
		}
		if err := tx.Where("applicant_id = ?", uid).Delete(&model.Application{}).Error; err != nil {
			return err
		}
		if err := tx.Where("user_id = ?", uid).Delete(&model.Case{}).Error; err != nil {
			return err
		}
		if err := tx.Delete(&model.User{}, "id = ?", uid).Error; err != nil {
			return err
		}
		return nil
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザーの削除に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "ユーザーを削除しました"})
}
