package case_handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	case_repository "github.com/kageyama0/chotto-rental/internal/repository/case"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// Create: 案件を作成します。
func (h *CaseHandler) Create(c *gin.Context) {
	var req CreateCaseRequest
	caseRepository := case_repository.NewCaseRepository(h.db)

	if err := c.ShouldBindJSON(&req); err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.JSON_PARSE_ERROR, err.Error())
		return
	}

	// パラメータの取得
	_, userID, errCode := util.GetParams(c, []string{})
	fmt.Println("userID0001", userID)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 案件の作成
	caseData := model.Case{
		UserID:         *userID,
		Title:          req.Title,
		Description:    req.Description,
		Category:       req.Category,
		Reward:         req.Reward,
		RequiredPeople: req.RequiredPeople,
		ScheduledDate:  req.ScheduledDate,
		StartTime:      req.StartTime,
		Duration:       req.Duration,
		Prefecture:     req.Prefecture,
		City:           req.City,
		Address:        req.Address,
		Status:         "open",
	}

	err := caseRepository.Create(&caseData)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	response := gin.H{
		"title":          req.Title,
		"description":    req.Description,
		"category":       req.Category,
		"reward":         req.Reward,
		"requiredPeople": req.RequiredPeople,
		"scheduledDate":  req.ScheduledDate,
		"startTime":      req.StartTime,
		"duration":       req.Duration,
		"prefecture":     req.Prefecture,
		"city":           req.City,
		"address":        req.Address,
		"status":         "open",
	}

	util.CreateResponse(c, http.StatusCreated, e.OK, response)
}
