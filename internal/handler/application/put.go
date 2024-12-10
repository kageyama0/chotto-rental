package application_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	application_repository "github.com/kageyama0/chotto-rental/internal/repository/application"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

type UpdateApplicationStatusRequest struct {
	Status string `json:"status" binding:"required,oneof=accepted rejected"`
}

// -- updateStatusParams: UpdateStatus関数で扱うパラメータが正しいかを確認し、正しい場合はそれらを返します。
func updateStatusParams(c *gin.Context) (applicationID *uuid.UUID, userID *uuid.UUID, errCode int) {
	var req UpdateApplicationStatusRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, nil, e.JSON_PARSE_ERROR
	}

	cParamsID := c.Param("id")
	applicationID, isValid := util.CheckUUID(c, cParamsID)
	if !isValid {
		return nil, nil, e.INVALID_ID
	}

	cUserID, _ := c.Get("userID")
	userID, isValid = util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, nil, e.INVALID_ID
	}

	return applicationID, userID, e.OK
}


// -- UpdateStatus; 応募のステータスを更新します。
func (h *ApplicationHandler) UpdateStatus(c *gin.Context) {
	var req UpdateApplicationStatusRequest

	applicationRepository := application_repository.NewApplicationRepository(h.db)
	applicationID, userID, errCode := updateStatusParams(c)
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
