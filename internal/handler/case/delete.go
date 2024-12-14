package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// -- Delete: 案件を削除します。
func (h *CaseHandler) Delete(c *gin.Context) {
	var caseData *model.Case
	caseRepository := case_repository.NewCaseRepository(h.db)

	// パラメータの取得
	params, userID, errCode := util.GetParams(c, []string{"case_id"})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}
	caseId := params["case_id"]

	// 案件の取得
	caseData, err := caseRepository.FindByID(caseId)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_CASE, nil)
		return
	}

	// ユーザーが案件の作成者であるかを確認
	if caseData.UserID != *userID {
		util.CreateResponse(c, http.StatusForbidden, e.FORBIDDEN_DELETE_CASE, nil)
	}

	//  案件の削除
	err = caseRepository.DeleteByID(caseId)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}
	util.CreateResponse(c, http.StatusNoContent, e.NO_CONTENT, nil)
}
