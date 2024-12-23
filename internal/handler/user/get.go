package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_repository "github.com/kageyama0/chotto-rental/internal/repository/user"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary ユーザーが自分のプロフィール情報を取得する
// @Description ユーザーIDに基づいて、ユーザーの基本情報を取得します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Success 200 {object} map[string]interface{} "ユーザー情報"
// @Failure 400 {object} util.Response "無効なユーザーID"
// @Failure 404 {object} util.Response "ユーザーが見つかりません"
// @Router /profile [get]
func (h *UserHandler) GetProfile(c *gin.Context) {
	userRepository := user_repository.NewUserRepository(h.db)

	_, userID, errCode := util.GetParams(c, []string{})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	user, errCode := userRepository.FindByID(*userID)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_USER, nil)
		return
	}

	util.CreateResponse(c, http.StatusOK, e.OK, util.StructToMap(user))
}
