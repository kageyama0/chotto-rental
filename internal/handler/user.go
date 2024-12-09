package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type UserHandler struct {
	db *gorm.DB
}

func NewUserHandler(db *gorm.DB) *UserHandler {
	return &UserHandler{db: db}
}

func (h *UserHandler) Get(c *gin.Context) {
	userID := c.Param("id")
	uid, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーID"})
		return
	}

	var user model.User
	if err := h.db.First(&user, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"email":        user.Email,
		"display_name": user.DisplayName,
		"trust_score":  user.TrustScore,
	})
}

type UpdateUserRequest struct {
	DisplayName string `json:"display_name" binding:"required"`
}

func (h *UserHandler) Update(c *gin.Context) {
	userID, _ := c.Get("userID")
	uid, _ := uuid.Parse(userID.(string))

	var req UpdateUserRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var user model.User
	if err := h.db.First(&user, "id = ?", uid).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}

	user.DisplayName = req.DisplayName
	if err := h.db.Save(&user).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "ユーザー情報の更新に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"id":           user.ID,
		"email":        user.Email,
		"display_name": user.DisplayName,
		"trust_score":  user.TrustScore,
	})
}

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

func (h *UserHandler) GetReviews(c *gin.Context) {
	userID := c.Param("id")
	uid, err := uuid.Parse(userID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "無効なユーザーID"})
		return
	}

	var reviews []model.Review
	if err := h.db.Preload("Reviewer").
		Where("reviewed_user_id = ?", uid).
		Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "レビューの取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
