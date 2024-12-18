package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

type CreateApplicationRequest struct {
	Message string `json:"message"`
}

// @Summary 応募作成
// @Description 案件への新規応募を作成します
// @Tags 応募
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param case_id path string true "案件ID"
// @Param request body CreateApplicationRequest true "応募メッセージ"
// @Success 201 {object} model.Application "応募作成成功"
// @Failure 400 {object} util.Response "無効なパラメータ、または案件が募集中でない"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 404 {object} util.Response "案件が見つかりません"
// @Failure 409 {object} util.Response "既に応募済み"
// @Failure 500 {object} util.Response "サーバーエラー"
func (h *ApplicationHandler) Create(c *gin.Context) {
	var req CreateApplicationRequest
	caseRepository := case_repository.NewCaseRepository(h.db)
	applicationRepository := application_repository.NewApplicationRepository(h.db)

	// パラメータの取得
	params, userID, errCode := util.GetParams(c, []string{"case_id"})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}
	caseID := params["case_id"]

	// 案件の取得
	var caseData *model.Case
	caseData, err := caseRepository.FindByID(caseID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_CASE, nil)
		return
	}

	// 案件が募集中か確認
	if caseData.Status != "open" {
		util.CreateResponse(c, http.StatusBadRequest, e.CASE_NOT_OPEN, nil)
	}

	// 既に応募していないか確認
	_, err = applicationRepository.FindByCaseIDAndApplicantID(caseID, *userID)
	if err == nil {
		util.CreateResponse(c, http.StatusConflict, e.ALREADY_APPLIED, nil)
		return
	}

	// 応募の作成
	application := model.Application{
		CaseID:      caseID,
		ApplicantID: *userID,
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
