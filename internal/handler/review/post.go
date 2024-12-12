package review_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	_ "github.com/kageyama0/chotto-rental/pkg/util"
)

type CreateReviewRequest struct {
	MatchingID     string `json:"matching_id" binding:"required"`
	ReviewedUserID string `json:"reviewed_user_id" binding:"required"`
	Score          int    `json:"score" binding:"required,min=1,max=5"`
	Comment        string `json:"comment"`
}

// @Summary レビュー作成
// @Description マッチング完了後、相手ユーザーへのレビューを作成します。1つのマッチングにつき1人のユーザーは1回のみレビュー可能です。
// @Tags レビュー
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param request body CreateReviewRequest true "レビュー情報"
// @Success 201 {object} model.Review "レビュー作成成功"
// @Failure 400 {object} util.Response "リクエストが不正です / 完了していないマッチングにはレビューできません / このマッチングに関係のないユーザーにレビューはできません"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 403 {object} util.Response "このマッチングにレビューを投稿する権限がありません"
// @Failure 404 {object} util.Response "マッチングが見つかりません"
// @Failure 409 {object} util.Response "既にレビューを投稿済みです"
// @Failure 500 {object} util.Response "サーバーエラー / 信頼スコアの更新に失敗しました"
// @Router /reviews [post]
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
