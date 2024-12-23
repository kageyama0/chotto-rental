package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	user_repository "github.com/kageyama0/chotto-rental/internal/repository/user"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

type UpdateUserRequest struct {
	DisplayName string `json:"displayName" binding:"required"`
}

// @Summary ユーザーのプロフィール情報を更新
// @Description ユーザーのプロフィール情報を更新します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer {token} 形式"
// @Param request body UpdateUserRequest true "更新するユーザー情報"
// @Success 200 {object} map[string]interface{} "更新後のユーザー情報"
// @Failure 400 {object} util.Response "リクエストが不正です"
// @Failure 401 {object} util.Response "認証エラー"
// @Failure 404 {object} util.Response "ユーザーが見つかりません"
// @Failure 500 {object} util.Response "サーバーエラー"
// @Router /profile [put]
func (h *UserHandler) UpdateProfile(c *gin.Context) {
	var req UpdateUserRequest
	userRepository := user_repository.NewUserRepository(h.db)

	_, userID, errCode := util.GetParams(c, []string{})
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, errCode, nil)
		return
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		util.CreateResponse(c, http.StatusBadRequest, e.JSON_PARSE_ERROR, nil)
		return
	}

	user, errCode := userRepository.FindByID(*userID)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusNotFound, e.NOT_FOUND_USER, nil)
		return
	}

	user.DisplayName = req.DisplayName
	errCode = userRepository.Update(*userID, user)
	if errCode != e.OK {
		util.CreateResponse(c, http.StatusInternalServerError, errCode, nil)
		return
	}

	util.CreateResponse(c, http.StatusOK, e.OK, util.StructToMap(user))
}
