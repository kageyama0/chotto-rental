package matching_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	matching_repository "github.com/kageyama0/chotto-rental/internal/repository/matching"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
	"gorm.io/gorm"
)


type CreateMatchingRequest struct {
	ApplicationID   string `json:"application_id" binding:"required"`
	MeetingLocation string `json:"meeting_location" binding:"required"`
}


// --CreateParams: Createのパラメータを取得する
func createParams(c *gin.Context) (applicationID *uuid.UUID, userID *uuid.UUID, errCode int) {
	var req CreateMatchingRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, nil, http.StatusBadRequest
	}

	applicationID, isValid := util.CheckUUID(c, req.ApplicationID)
	if !isValid {
		return nil, nil, http.StatusBadRequest
	}

	cUserID, _ := c.Get("userID")
	userID, isValid = util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, nil, http.StatusBadRequest
	}

	return applicationID, userID, http.StatusOK
}

// --Create: マッチングを作成する
func (h *MatchingHandler) Create(c *gin.Context) {
	var req CreateMatchingRequest
	var application Application
	applicationRepository := application_repository.NewApplicationRepository(h.db)

	// パラメータの取得
	applicationID, userID, errCode := createParams(c)
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
