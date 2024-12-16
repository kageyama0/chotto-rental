package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary 案件詳細取得
// @Description 指定されたIDの案件情報を取得します
// @Tags 案件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param id path string true "案件ID"
// @Success 200 {object} util.Response{data=map[string]model.Case} "案件情報"
// @Failure 400 {object} util.Response "リクエストが不正です"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 404 {object} util.Response "案件が見つかりません"
// @Router /cases/{id} [get]
func (h *CaseHandler) Get(c *gin.Context) {
	var caseData *model.Case
	caseRepository := case_repository.NewCaseRepository(h.db)

	// パラメータの取得
	params, _, errCode := util.GetParams(c, []string{"case_id"})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}
	caseID := params["case_id"]

	// 案件の取得
	caseData, err := caseRepository.FindByID(caseID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_CASE, nil)
		return
	}
	util.CreateResponse(c, http.StatusOK, e.OK, map[string]interface{}{
		"case": caseData,
	})
}
