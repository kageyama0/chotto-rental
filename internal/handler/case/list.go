package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary 案件一覧取得
// @Description 全ての案件情報を取得します
// @Tags 案件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Success 200 {object} util.Response{data=map[string][]model.Case} "案件一覧"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /cases [get]
// TODO: Paginationを追加する
func (h *CaseHandler) List(c *gin.Context) {
	var cases []model.Case
	caseRepository := case_repository.NewCaseRepository(h.db)

	// 案件の取得
	cases, err := caseRepository.FindAll()
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	util.CreateResponse(c, http.StatusOK, e.OK, map[string]interface{}{
		"cases": cases,
	})
}
