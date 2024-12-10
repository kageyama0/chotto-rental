package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// -- listParams
// List関数で扱うパラメータが正しいかを確認し、正しい場合はそれらを返します。
func listParams(c *gin.Context) (userID *uuid.UUID, errCode int) {
	// ユーザーIDの取得
	cUserID, _ := c.Get("userID")
	userID, isValid := util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, e.INVALID_ID
	}

	return userID, e.OK
}


// -- List
// ユーザーが応募している全ての案件を取得します。
func (h *ApplicationHandler) List(c *gin.Context) {
	applicationRepository := application_repository.NewApplicationRepository(h.db)

	userID, errCode := listParams(c)
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
