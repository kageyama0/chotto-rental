package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

func (h *AuthHandler) Refresh(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		util.CreateResponse(c, http.StatusUnauthorized, e.INVALID_SESSION_ID, nil)
		return
	}

	deviceInfo := DeviceInfo{
		UserAgent: c.GetHeader("User-Agent"),
		IP:        c.ClientIP(),
	}

	parsedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.INVALID_SESSION_ID, nil)
		return
	}

	statusCode, errCode := h.authService.Refresh(c, parsedSessionID, deviceInfo)
	if errCode != e.OK {
		util.CreateResponse(c, statusCode, errCode, nil)
	}

	util.CreateResponse(c, http.StatusOK, e.OK, gin.H{})
}
