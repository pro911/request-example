package util

import (
	"github.com/gin-gonic/gin"
	"github.com/pro911/request-example/pkg/e"
)

func Success(data interface{}, msg string) map[string]interface{} {
	if data == nil {
		data = make(map[string]interface{})
	}
	if msg == "" {
		msg = e.GetMsg(e.SUCCESS)
	}
	return gin.H{
		"code": e.SUCCESS,
		"msg":  msg,
		"data": data,
	}
}

func ErrorFail(code int, msg string, data interface{}) map[string]interface{} {
	if msg == "" {
		msg = e.GetMsg(code)
	}
	if data == nil {
		data = make(map[string]interface{})
	}
	return gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	}
}
