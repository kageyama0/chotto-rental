package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)


type CreateApplicationRequest struct {
	CaseID  string `json:"case_id" binding:"required"`
	Message string `json:"message"`
}


// -- createStatusParams
// CreateStatus関数で扱うパラメータが正しいかを確認し、正しい場合はそれらを返します。
// パラメータが正しくない場合は、エラーコードを返します。
func createStatusParams(c *gin.Context) (caseID *uuid.UUID, userID *uuid.UUID, errCode int) {
	var req CreateApplicationRequest

	// リクエストのパース
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, nil, e.JSON_PARSE_ERROR
	}

	// パラメータの取得
	cCaseID := req.CaseID
	caseID, isValid := util.CheckUUID(c, cCaseID)
	if !isValid {
		return nil, nil, e.INVALID_ID
	}

	// ユーザーIDの取得
	cUserID, _ := c.Get("userID")
	userID, isValid = util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, nil, e.INVALID_ID
	}

	return caseID, userID, e.OK
}


// -- Create
// 案件に対する応募を作成します。
func (h *ApplicationHandler) Create(c *gin.Context) {
	var req CreateApplicationRequest
	caseRepository := case_repository.NewCaseRepository(h.db)
	applicationRepository := application_repository.NewApplicationRepository(h.db)

	// パラメータの取得
	caseID, uid, errCode := createStatusParams(c)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 案件の取得
	var caseData *model.Case
	caseData, err := caseRepository.FindByID(*caseID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_CASE, nil)
		return
	}

	// 案件が募集中か確認
	if caseData.Status != "open" {
		util.CreateResponse(c, http.StatusBadRequest, e.CASE_NOT_OPEN, nil)
	}

	// 既に応募していないか確認
	_, err = applicationRepository.FindByCaseIDAndApplicantID(*caseID, *uid)
	if err == nil {
		util.CreateResponse(c, http.StatusConflict, e.ALREADY_APPLIED, nil)
		return
	}

	// 応募の作成
	application := model.Application{
		CaseID:      *caseID,
		ApplicantID: *uid,
		Message:     req.Message,
		Status:      "pending",
	}
	err = applicationRepository.Create(&application)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	util.CreateResponse(c, http.StatusCreated, e.CREATED, application)
}
