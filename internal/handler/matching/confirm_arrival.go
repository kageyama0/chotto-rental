package matching_handler

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/kageyama0/chotto-rental/internal/model"
	matching_repository "github.com/kageyama0/chotto-rental/internal/repository/matching"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary 到着確認
// @Description マッチング成立後の到着確認を行います。依頼者とヘルパー両方の確認が完了すると、マッチングが完了状態になります。
// @Tags マッチング
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param id path string true "マッチングID"
// @Success 200 {object} util.Response{data=model.Matching} "OK"
// @Failure 400 {object} util.Response "リクエストが不正です / この案件は募集を終了しています / 確認期限が過ぎています"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 403 {object} util.Response "この操作を行う権限がありません"
// @Failure 404 {object} util.Response "Not Found"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /matchings/{id}/confirm-arrival [post]
func (h *MatchingHandler) ConfirmArrival(c *gin.Context) {
	var matching Matching
	matchingRepository := matching_repository.NewMatchingRepository(h.db)

	// パラメータの取得
	params, userID, errCode := util.GetParams(c, []string{"matching_id"})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	// マッチングの取得
	matching, err := matchingRepository.FindByID(params["matching_id"])
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
