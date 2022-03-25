package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/pro911/request-example/middleware/jwt"
	"github.com/pro911/request-example/pkg/setting"
	"github.com/pro911/request-example/routers/demo"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	gin.SetMode(setting.RunMode)

	r.GET("/test", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "test",
		})
	})

	apiv1 := r.Group("/api")
	{
		//登录
		apiv1.POST("/demo/login", demo.Login)
		apidemo := apiv1.Group("/demo")
		apidemo.Use(jwt.JWT())
		{
			//添加新闻
			apidemo.POST("/news", demo.AddNew)
			//新闻列表
			apidemo.GET("/news_list", demo.NewsList)
			//新闻详情
			apidemo.GET("/news_details", demo.NewsDetails)
			//新闻评论
			apidemo.POST("/news_comment", demo.NewsComment)
			//收藏新闻
			apidemo.POST("/collect_news", demo.CollectNews)
			//删除评论
			apidemo.DELETE("/delete_comment", demo.DeleteComment)
		}
	}

	return r
}
