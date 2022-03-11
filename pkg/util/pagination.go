package util

import (
	"github.com/gin-gonic/gin"
	"github.com/pro911/request-example/pkg/setting"
	"github.com/unknwon/com"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	pageSize, _ := com.StrTo(c.DefaultQuery("page_size", string(setting.PageSize))).Int()
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
