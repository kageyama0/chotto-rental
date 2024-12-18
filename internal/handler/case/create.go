package case_handler

import (
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

	// パラメータの取得
	_, userID, errCode := util.GetParams(c, []string{})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// 案件の作成
	caseData := model.Case{
		UserID:          *userID,
		Title:           req.Title,
		Description:     req.Description,
		Reward:          req.Reward,
		Location:        req.Location,
		ScheduledDate:   req.ScheduledDate,
		DurationMinutes: req.DurationMinutes,
		Status:          "open",
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
