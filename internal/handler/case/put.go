package case_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary 案件更新
// @Description 指定された案件の情報を更新します（案件作成者のみ可能）
// @Tags 案件
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param id path string true "案件ID"
// @Param request body CreateCaseRequest true "更新する案件情報"
// @Success 200 {object} model.Case "OK"
// @Failure 400 {object} util.Response "リクエストが不正です"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 403 {object} util.Response "この操作を行う権限がありません"
// @Failure 404 {object} util.Response "案件が見つかりません"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /cases/{id} [put]
func (h *CaseHandler) Update(c *gin.Context) {
	id := c.Param("id")
	var req CreateCaseRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}

	var caseData model.Case
	if err := h.db.First(&caseData, "id = ?", id).Error; err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND, nil)
		return
	}

	userID, _ := c.Get("userID")
	if caseData.UserID.String() != userID.(string) {
		util.CreateResponse(c, http.StatusForbidden, e.FORBIDDEN, nil)
		return
	}

	caseData.Title = req.Title
	caseData.Description = req.Description
	caseData.Reward = req.Reward
	caseData.Location = req.Location
	caseData.ScheduledDate = req.ScheduledDate
	caseData.DurationMinutes = req.DurationMinutes

	if err := h.db.Save(&caseData).Error; err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	c.JSON(http.StatusOK, caseData)
}
