package matching_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/kageyama0/chotto-rental/internal/model"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	matching_repository "github.com/kageyama0/chotto-rental/internal/repository/matching"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
	"gorm.io/gorm"
)

type CreateMatchingRequest struct {
	MeetingLocation string `json:"meetingLocation" binding:"required"`
}

// @Summary マッチング作成
// @Description 応募を承認してマッチングを作成します。応募のステータスを「accepted」に、案件のステータスを「matched」に更新します。
// @Tags マッチング
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param application_id path string true "応募ID"
// @Param request body CreateMatchingRequest true "マッチング作成情報"
// @Success 201 {object} util.Response{data=model.Matching} "Created"
// @Failure 400 {object} util.Response "リクエストが不正です"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 403 {object} util.Response "この操作を行う権限がありません"
// @Failure 404 {object} util.Response "案件が見つかりません"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /applications/{application_id}/matching [post]
func (h *MatchingHandler) Create(c *gin.Context) {
	var req CreateMatchingRequest
	var application Application
	applicationRepository := application_repository.NewApplicationRepository(h.db)

	// パラメータの取得
	params, userID, errCode := util.GetParams(c, []string{"application_id"})
	applicationID := params["application_id"]
	if errCode != http.StatusOK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 応募の取得
	application, err := applicationRepository.FindByIDWithCase(applicationID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_CASE, nil)
		return
	}

	// ユーザーの権限チェック
	if application.Case.UserID != *userID {
		util.CreateResponse(c, http.StatusForbidden, e.FORBIDDEN, nil)
		return
	}

	matching := Matching{
		CaseID:                      application.CaseID,
		RequesterID:                 application.Case.UserID,
		HelperID:                    application.ApplicantID,
		MeetingLocation:             req.MeetingLocation,
		ArrivalConfirmationDeadline: application.Case.ScheduledDate.Add(15 * time.Minute),
		Status:                      "active",
	}

	// トランザクション開始
	if err := h.db.Transaction(func(tx *gorm.DB) error {
		matchingRepository := matching_repository.NewMatchingRepository(tx)
		applicationRepository := application_repository.NewApplicationRepository(tx)
		caseRepository := case_repository.NewCaseRepository(tx)

		// マッチングの作成
		if err := matchingRepository.Create(&matching); err != nil {
			return err
		}

		// 応募のステータスを「accepted」に更新
		application.Status = "accepted"
		err := applicationRepository.Update(&application)
		if err != nil {
			return err
		}

		// 案件のステータスを「matched」に更新
		err = caseRepository.UpdateStatus(application.CaseID, "matched")
		if err != nil {
			return err
		}
		return nil
	}); err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	util.CreateResponse(c, http.StatusCreated, e.OK, matching)
}
