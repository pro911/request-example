package util

import (
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"request-example/config"
	"strconv"
)

func GetPage(c *gin.Context) int {
	result := 0
	page, _ := com.StrTo(c.Query("page")).Int()
	pageSize, _ := com.StrTo(c.DefaultQuery("page_size", strconv.Itoa(config.AppConf.PageSize))).Int()
	if page > 0 {
		result = (page - 1) * pageSize
	}
	return result
}
