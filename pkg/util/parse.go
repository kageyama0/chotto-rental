package util

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// uuidとして有効かどうかをチェックする
// 有効だった場合: parseしたuuidとtrueを返す
// 無効だった場合: nilとfalseを返す
func CheckUUID(c *gin.Context, uid string) (*uuid.UUID, bool) {
	checked_uuid, err := uuid.Parse(uid)
	if err != nil {
		return nil, false
	}

	return &checked_uuid, true
}
