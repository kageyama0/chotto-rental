package util

import (
	"github.com/gin-gonic/gin"

	"github.com/kageyama0/chotto-rental/pkg/e"
)


type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func CreateResponse(c *gin.Context, httpCode, errCode int, data interface{}) {
	c.JSON(httpCode, Response{
		Code: errCode,
		Msg:  e.GetMsg(errCode),
		Data: data,
	})
}
