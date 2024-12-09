package handler

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	"gorm.io/gorm"
)

type ReviewHandler struct {
	db *gorm.DB
}

func NewReviewHandler(db *gorm.DB) *ReviewHandler {
	return &ReviewHandler{db: db}
}

type CreateReviewRequest struct {
	MatchingID     string `json:"matching_id" binding:"required"`
	ReviewedUserID string `json:"reviewed_user_id" binding:"required"`
	Score          int    `json:"score" binding:"required,min=1,max=5"`
	Comment        string `json:"comment"`
}

func (h *ReviewHandler) Create(c *gin.Context) {
	var req CreateReviewRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := c.Get("userID")
	reviewerID, _ := uuid.Parse(userID.(string))
	reviewedUserID, _ := uuid.Parse(req.ReviewedUserID)
	matchingID, _ := uuid.Parse(req.MatchingID)

	// マッチングの存在確認と権限チェック
	var matching model.Matching
	if err := h.db.First(&matching, "id = ?", matchingID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "マッチングが見つかりません"})
		return
	}

	if matching.Status != "completed" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "完了していないマッチングにはレビューできません"})
		return
	}

	if matching.RequesterID != reviewerID && matching.HelperID != reviewerID {
		c.JSON(http.StatusForbidden, gin.H{"error": "このマッチングにレビューを投稿する権限がありません"})
		return
	}

	if matching.RequesterID != reviewedUserID && matching.HelperID != reviewedUserID {
		c.JSON(http.StatusBadRequest, gin.H{"error": "このマッチングに関係のないユーザーにレビューはできません"})
		return
	}

	// 既存レビューの確認
	var existingReview model.Review
	if err := h.db.Where("matching_id = ? AND reviewer_id = ?", matchingID, reviewerID).First(&existingReview).Error; err == nil {
		c.JSON(http.StatusConflict, gin.H{"error": "既にレビューを投稿済みです"})
		return
	}

	review := model.Review{
		MatchingID:     matchingID,
		ReviewerID:     reviewerID,
		ReviewedUserID: reviewedUserID,
		Score:          req.Score,
		Comment:        req.Comment,
	}

	if err := h.db.Create(&review).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "レビューの作成に失敗しました"})
		return
	}

	// ユーザーの信頼スコアを更新
	if err := h.updateUserTrustScore(reviewedUserID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "信頼スコアの更新に失敗しました"})
		return
	}

	c.JSON(http.StatusCreated, review)
}

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

func (h *ReviewHandler) List(c *gin.Context) {
	userID := c.Query("user_id")
	var reviews []model.Review

	query := h.db.Preload("Reviewer").Preload("ReviewedUser").Preload("Matching")
	if userID != "" {
		query = query.Where("reviewed_user_id = ?", userID)
	}

	if err := query.Find(&reviews).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "レビューの取得に失敗しました"})
		return
	}

	c.JSON(http.StatusOK, reviews)
}
