package user_handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/internal/model"
	"github.com/kageyama0/chotto-rental/pkg/e"
	"github.com/kageyama0/chotto-rental/pkg/util"
)

// @Summary ユーザー情報取得
// @Description ユーザーIDに基づいてユーザーの基本情報を取得します
// @Tags ユーザー
// @Accept json
// @Produce json
// @Param id path string true "ユーザーID"
// @Success 200 {object} map[string]interface{} "ユーザー情報"
// @Failure 400 {object} util.Response "無効なユーザーID"
// @Failure 404 {object} util.Response "ユーザーが見つかりません"
// @Router /users/{id} [get]
func (h *UserHandler) Get(c *gin.Context) {
	_, userID, err := util.GetParams(c, []string{})
	if err != e.OK {
		util.CreateResponse(c, http.StatusBadRequest, err, nil)
		return
	}

	var user model.User
	if err := h.db.First(&user, "id = ?", userID).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "ユーザーが見つかりません"})
		return
	}
	
	util.CreateResponse(c, http.StatusOK, e.OK, map[string]interface{}{
		"user": user,
	})
}
