package auth_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

func (h *AuthHandler) Logout(c *gin.Context) {
	sessionID, err := c.Cookie("session_id")
	if err != nil {
		util.CreateResponse(c, http.StatusUnauthorized, e.NOT_FOUND_SESSION, nil)
		return
	}

	parsedSessionID, err := uuid.Parse(sessionID)
	if err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.INVALID_SESSION_ID, nil)
		return
	}

	statusCode, errCode := h.authService.Logout(c, parsedSessionID)
	if errCode != e.OK {
		util.CreateResponse(c, statusCode, errCode, nil)
		return
	}

	// クッキーの削除
	c.SetCookie(
		"session_id",
		"",
		-1,
		"/",
		"",
		true,
		true,
	)

	util.CreateResponse(c, http.StatusOK, e.OK, nil)
}
