package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// -- GetParams: Get関数で扱うパラメータが正しいかを確認し、正しい場合はそれらを返します。
func getParams(c *gin.Context) (caseID *uuid.UUID, errCode int) {
	cCaseID := c.Param("id")
	caseID, isValid := util.CheckUUID(c, cCaseID)
	if !isValid {
		return nil, e.INVALID_ID
	}

	return caseID, e.OK
}

// -- Get: 指定されたIDの案件を取得します。
func (h *CaseHandler) Get(c *gin.Context) {
	var caseData *model.Case
	caseRepository := case_repository.NewCaseRepository(h.db)

	// パラメータの取得
	caseID, errCode := getParams(c)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}


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

// -- List: 全ての案件を取得します。
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
