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

// -- deleteParams: Delete関数で扱うパラメータが正しいかを確認し、正しい場合はそれらを返します。
func deleteParams(c *gin.Context) (caseID *uuid.UUID, userID *uuid.UUID, errCode int) {
	// パラメータの取得
	cCaseID := c.Param("id")
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


// -- Delete: 案件を削除します。
func (h *CaseHandler) Delete(c *gin.Context) {
	var caseData *model.Case
	caseRepository := case_repository.NewCaseRepository(h.db)

	// パラメータの取得
	caseId, userID, errCode := deleteParams(c)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

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
