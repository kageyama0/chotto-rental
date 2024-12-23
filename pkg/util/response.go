package util

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kageyama0/chotto-rental/pkg/e"
)

// @Description HTTPレスポンス
type Response struct {
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CreateResponse(c *gin.Context, httpCode int, msgCode int, data interface{}) {
	log.Printf("httpCode: %d, msgCode: %d, data: %v", httpCode, msgCode, data)

	c.JSON(httpCode, Response{
		Msg:  e.GetMsg(msgCode),
		Data: data,
	})
}
