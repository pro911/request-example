package demo

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/pro911/request-example/models"
	"github.com/pro911/request-example/pkg/e"
	"github.com/pro911/request-example/pkg/util"
	"github.com/unknwon/com"
	"net/http"
	"strconv"
	"time"
)

// Login 登录
func Login(c *gin.Context) {
	mobile := c.DefaultPostForm("mobile", "")
	verCode := com.StrTo(c.DefaultPostForm("ver_code", "0")).MustInt()

	valid := validation.Validation{}

	valid.Required(mobile, "mobile").Message("手机号不能为空")
	valid.Mobile(mobile, "mobile").Message("手机号格式不正确")
	valid.Required(verCode, "ver_code").Message("验证码不能为空")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
			return
		}
	}

	data := make(map[string]interface{})

	unixNano := time.Now().UnixNano()

	bytes := []byte(strconv.Itoa(int(unixNano)))

	data["token"] = fmt.Sprintf("%x", md5.Sum(bytes))

	c.JSON(http.StatusOK, util.Success(data))
}

type AddNewJson struct {
	Title   string `json:"title" binding:"required"`
	Author  string `json:"author" binding:"required"`
	Url     string `json:"url" binding:"required"`
	Content string `json:"content" binding:"required"`
}

func AddNew(c *gin.Context) {
	// 声明接收的变量
	var json AddNewJson
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, util.ErrorFail(e.ERROR_ADD_FAIL, err.Error(), nil))
		return
	}
	fmt.Println(json)

	valid := validation.Validation{}
	valid.Required(json.Title, "title").Message("标题不能为空")
	valid.MaxSize(json.Title, 100, "title").Message("标题最长为100字符")
	valid.Required(json.Author, "author").Message("作者不能为空")
	valid.MaxSize(json.Author, 50, "author").Message("作者最长为50字符")
	valid.Required(json.Url, "url").Message("链接地址不能为空")
	valid.Required(json.Content, "content").Message("内容不能为空")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
			return
		}
	}

	if !models.AddNew(json.Title, json.Author, json.Url, json.Content) {
		c.JSON(http.StatusOK, util.ErrorFail(e.ERROR_ADD_FAIL, "", nil))
		return
	}
	c.JSON(http.StatusOK, util.Success(make(map[string]interface{})))
	return
}

// NewsList 新闻列表
func NewsList(c *gin.Context) {
	mobile := c.DefaultQuery("mobile", "")
	themeNews := c.DefaultQuery("theme_news", "")
	page := com.StrTo(c.DefaultQuery("page", "1")).MustInt()
	pageSize := com.StrTo(c.DefaultQuery("page_size", "20")).MustInt()

	valid := validation.Validation{}
	valid.Required(mobile, "mobile").Message("手机号不能为空")
	valid.Mobile(mobile, "mobile").Message("手机号格式不正确")
	valid.Required(themeNews, "theme_news").Message("新闻专题不能为空")
	valid.Required(page, "page").Message("page字段不能为空")
	valid.Required(pageSize, "page_size").Message("page_size字段不能为空")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
			return
		}
	}

	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	data["lists"] = models.GetNews(util.GetPage(c), pageSize, maps)
	data["total"] = models.GetNewTotal(maps)

	c.JSON(http.StatusOK, util.Success(data))
}

// NewsDetails 新闻详情
func NewsDetails(c *gin.Context) {

}

// NewsComment 新闻评论
func NewsComment(c *gin.Context) {

}

// CollectNews 收藏新闻
func CollectNews(c *gin.Context) {

}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {

}
