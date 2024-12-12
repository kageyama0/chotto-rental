package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	matching_repository "github.com/kageyama0/chotto-rental/internal/repository/matching"
	review_repository "github.com/kageyama0/chotto-rental/internal/repository/review"
	user_repository "github.com/kageyama0/chotto-rental/internal/repository/user"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
	"gorm.io/gorm"
)

// -- DeleteParams: Deleteのパラメータを取得する
func deleteParams(c *gin.Context) (userID *uuid.UUID, errCode int) {
	cUserID, _ := c.Get("userID")
	userID, isValid := util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, http.StatusBadRequest
	}

	return userID, http.StatusOK
}

// @Summary ユーザー削除
// @Description ユーザーと関連する全てのデータ（レビュー、マッチング、応募、案件）を削除します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Success 204 {object} util.Response "No Content"
// @Failure 400 {object} util.Response "リクエストが不正です"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /users/me [delete]
func (h *UserHandler) Delete(c *gin.Context) {
	// パラメータの取得
	userID, errCode := deleteParams(c)
	if errCode != http.StatusOK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 関連データの削除をするため、トランザクション開始
	if err := h.db.Transaction(func(tx *gorm.DB) error {
		applicationRepository := application_repository.NewApplicationRepository(tx)
		caseRepository := case_repository.NewCaseRepository(tx)
		matchingRepository := matching_repository.NewMatchingRepository(tx)
		reviewRepository := review_repository.NewReviewRepository(tx)
		userRepository := user_repository.NewUserRepository(tx)

		// レビューの削除
		err := reviewRepository.DeleteByReviewerIDAndReviewedID(userID)
		if err != nil {
			return err
		}

		// マッチングの削除
		err = matchingRepository.DeleteByRequesterIDAndHelperID(userID)
		if err != nil {
			return err
		}

		// アプリケーションの削除
		err = applicationRepository.DeleteByUserID(userID)
		if err != nil {
			return err
		}

		// 案件の削除
		err = caseRepository.DeleteByUserID(userID)
		if err != nil {
			return err
		}

		// ユーザーの削除
		err = userRepository.DeleteByID(userID)
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, http.StatusInternalServerError, nil)
		return
	}

	util.CreateResponse(c, http.StatusNoContent, e.NO_CONTENT, nil)
}
