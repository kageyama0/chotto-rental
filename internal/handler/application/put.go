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

func (h *ApplicationHandler) UpdateStatus(c *gin.Context) {
	var req UpdateApplicationStatusRequest

	applicationRepository := application_repository.NewApplicationRepository(h.db)
	applicationID, userID, errCode := updateStatusParams(c)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	application, err := applicationRepository.FindByIDWithCase(applicationID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_APPLICATION, nil)
		return
	}

	if application.Case.UserID != *userID {
		util.CreateResponse(c, http.StatusForbidden, e.FORBIDDEN_UPDATE_CASE, nil)
		return
	}

	application.Status = req.Status
	if err := h.db.Save(&application).Error; err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	c.JSON(http.StatusOK, application)
}
