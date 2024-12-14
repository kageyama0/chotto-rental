package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/kageyama0/chotto-rental/internal/model"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary 応募一覧取得
// @Description ユーザーの全ての応募履歴を案件情報と共に取得します
// @Tags 応募
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Success 200 {object} []model.Application "応募一覧の取得成功"
// @Failure 400 {object} util.Response "無効なパラメータ"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /applications [get]
func (h *ApplicationHandler) List(c *gin.Context) {
	applicationRepository := application_repository.NewApplicationRepository(h.db)

	_, userID, errCode := util.GetParams(c, []string{})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	applications, err := applicationRepository.FindAllByIDWithCase(userID)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	util.CreateResponse(c, http.StatusOK, e.OK, map[string]interface{}{
		"applications": applications,
	})
}
