package case_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kageyama0/chotto-rental/internal/model"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

type CreateCaseRequest struct {
	Title          string    `json:"title" binding:"required"`
	Description    string    `json:"description" binding:"required"`
	Reward         int       `json:"reward" binding:"required,min=0"`
	Location       string    `json:"location" binding:"required"`
	ScheduledDate  time.Time   `json:"scheduled_date" binding:"required"`
	DurationMinutes int      `json:"duration_minutes" binding:"required,min=1"`
}


// -- createParams: Create関数で扱うパラメータが正しいかを確認し、正しい場合はそれらを返します。
func createParams(c *gin.Context) (userID *uuid.UUID, errCode int) {
	var req CreateCaseRequest

	// リクエストのパース
	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, e.JSON_PARSE_ERROR
	}

	// ユーザーIDの取得
	cUserID, _ := c.Get("userID")
	userID, isValid := util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, e.INVALID_ID
	}

	return userID, e.OK
}


// Create: 案件を作成します。
func (h *CaseHandler) Create(c *gin.Context) {
	var req CreateCaseRequest
	caseRepository := case_repository.NewCaseRepository(h.db)

	// パラメータの取得
	userID, errCode := createParams(c)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 案件の作成
	caseData := model.Case{
		UserID:         *userID,
		Title:          req.Title,
		Description:    req.Description,
		Reward:         req.Reward,
		Location:       req.Location,
		ScheduledDate:  req.ScheduledDate,
		DurationMinutes: req.DurationMinutes,
		Status:         "open",
	}
	err := caseRepository.Create(&caseData)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	util.CreateResponse(c, http.StatusCreated, e.OK, map[string]interface{}{
		"case": caseData,
	})
}
