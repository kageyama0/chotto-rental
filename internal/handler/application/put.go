package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/kageyama0/chotto-rental/internal/model"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

type UpdateApplicationStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=accepted rejected"`
}

// @Summary 応募ステータス更新
// @Description 応募のステータスを更新します（案件作成者のみ可能）
// @Tags 応募
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param application_id path string true "応募ID"
// @Param request body UpdateApplicationStatusRequest true "更新するステータス情報"
// @Success 200 {object} model.Application  "OK"
// @Failure 400 {object} util.Response "無効なパラメータ"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 403 {object} util.Response "この応募のステータスを更新する権限がありません",
// @Failure 404 {object} util.Response "応募が見つかりません"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /applications/{application_id}/status [put]
func (h *ApplicationHandler) UpdateStatus(c *gin.Context) {
	var req UpdateApplicationStatusRequest

	applicationRepository := application_repository.NewApplicationRepository(h.db)
	params, userID, errCode := util.GetParams(c, []string{"application_id"})
	applicationID := params["application_id"]
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 応募の取得
	application, err := applicationRepository.FindByIDWithCase(applicationID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_APPLICATION, nil)
		return
	}

	// 案件の作成者であるかを確認
	if application.Case.UserID != *userID {
		util.CreateResponse(c, http.StatusForbidden, e.FORBIDDEN_UPDATE_APPLICATION, nil)
		return
	}

	// ステータスの更新
	application.Status = req.Status
	err = applicationRepository.Update(&application)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	c.JSON(http.StatusOK, application)
}
