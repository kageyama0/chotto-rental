package matching_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	matching_repository "github.com/kageyama0/chotto-rental/internal/repository/matching"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// --confirmArrivalParams: ConfirmArrivalのパラメータを取得する
func confirmArrivalParams(c *gin.Context) (matchingID *uuid.UUID, userID *uuid.UUID, errCode int) {
	cParamID:= c.Param("id")
	matchingID, isValid := util.CheckUUID(c, cParamID)
	if !isValid {
		return nil, nil, e.INVALID_ID
	}

	cUserID, _ := c.Get("userID")
	userID, isValid = util.CheckUUID(c, cUserID.(string))
	if !isValid {
		return nil, nil, e.INVALID_ID
	}

	return matchingID, userID, e.OK
}

// -- ConfirmArrival: 到着確認を行う
func (h *MatchingHandler) ConfirmArrival(c *gin.Context) {
	var matching Matching
	matchingRepository := matching_repository.NewMatchingRepository(h.db)

	// パラメータの取得
	matchingID, userID, errCode := confirmArrivalParams(c)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// マッチングの取得
	matching, err := matchingRepository.FindByID(matchingID)
	if err != nil {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND, nil)
		return
	}

	// マッチングの確認
	if matching.Status != "active" {
		util.CreateResponse(c, http.StatusBadRequest, e.BAD_REQUEST, nil)
		return
	}

	// 到着確認の確認期限が過ぎているかどうかをチェック
	if time.Now().After(matching.ArrivalConfirmationDeadline) {
		util.CreateResponse(c, http.StatusBadRequest, e.OVER_CONFIRMATION_DEADLINE, nil)
		return
	}

	// それぞれのユーザーの到着確認フラグを更新する
	switch *userID {
	case matching.RequesterID:
		matching.ArrivalConfirmedByRequester = true
	case matching.HelperID:
		matching.ArrivalConfirmedByHelper = true
	default:
		util.CreateResponse(c, http.StatusForbidden, e.FORBIDDEN, nil)
		return
	}

	// どちらも確認済みの場合はマッチングを完了状態にする
	if matching.ArrivalConfirmedByRequester && matching.ArrivalConfirmedByHelper {
		matching.Status = "completed"
	}

	err = matchingRepository.Update(&matching)
	if err != nil {
		util.CreateResponse(c, http.StatusInternalServerError, e.SERVER_ERROR, nil)
		return
	}

	util.CreateResponse(c, http.StatusOK, e.OK, matching)
}
