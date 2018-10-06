package handler

import (
	"net/http"

	"apiserver/demo/pkg/errno"

	"github.com/gin-gonic/gin"
)

/*
返回消息使用统一入口
*/
type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, message := errno.DecodeErr(err)

	// always return http.StatusOK
	c.JSON(http.StatusOK, Response{
		Code:    code,
		Message: message,
		Data:    data,
	})
}
