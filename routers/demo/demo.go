package demo

import (
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
	"math"
	"math/rand"
	"net/http"
	"request-example/models"
	"request-example/pkg/e"
	"request-example/pkg/util"
)

type LoginJson struct {
	Mobile  string `json:"mobile" valid:"Required;MaxSize(11)"`
	VerCode string `json:"ver_code" valid:"Required;MaxSize(4)"`
}

// Login 登录
func Login(c *gin.Context) {

	//声明接收的变量
	var json LoginJson

	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, util.ErrorFail(e.ERROR_ADD_FAIL, err.Error(), nil))
		return
	}

	//mobile := c.DefaultPostForm("mobile", "")
	//verCode := com.StrTo(c.DefaultPostForm("ver_code", "0")).MustInt()

	valid := validation.Validation{}
	valid.Valid(&json)

	valid.Required(json.Mobile, "mobile").Message("手机号不能为空")
	valid.Mobile(json.Mobile, "mobile").Message("手机号格式不正确")
	valid.Required(json.Mobile, "ver_code").Message("验证码不能为空")

	if valid.HasErrors() {
		for _, err := range valid.Errors {
			c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
			return
		}
	}

	data := make(map[string]interface{})

	//unixNano := time.Now().UnixNano()
	//
	//bytes := []byte(strconv.Itoa(int(unixNano)))
	//
	//data["token"] = fmt.Sprintf("%x", md5.Sum(bytes))

	token, err := util.GenerateToken(json.Mobile, json.VerCode)
	if err != nil {
		c.JSON(http.StatusOK, util.ErrorFail(e.ERROR_AUTH_TOKEN, "", nil))
		return
	}
	data["token"] = token
	c.JSON(http.StatusOK, util.Success(data, ""))
	return
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
	c.JSON(http.StatusOK, util.Success(make(map[string]interface{}), "新增成功"))
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
	data["page_size"] = pageSize
	data["cur_page"] = page
	data["last_page"] = int(math.Ceil(float64(models.GetNewTotal(maps)) / float64(pageSize)))

	c.JSON(http.StatusOK, util.Success(data, ""))
}

// NewsDetails 新闻详情
func NewsDetails(c *gin.Context) {
	id := com.StrTo(c.DefaultQuery("id", "0")).MustInt()
	if id <= 0 {
		c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, "", nil))
		return
	}
	new := models.GetNew(id)
	if new.ID <= 0 {
		c.JSON(http.StatusOK, util.ErrorFail(e.ERROR_NOT_EXIST_ARTICLE, "", nil))
		return
	}
	c.JSON(http.StatusOK, util.Success(new, ""))
	return
}

type NewsCommentJson struct {
	ID      int    `json:"id" binding:"required"`
	Content string `json:"content" binding:"required"`
}

// NewsComment 新闻评论
func NewsComment(c *gin.Context) {
	// 声明接收的变量
	var json NewsCommentJson
	// 将request的body中的数据，自动按照json格式解析到结构体
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, util.ErrorFail(e.ERROR_ADD_FAIL, err.Error(), nil))
		return
	}

	valid := validation.Validation{}
	valid.Min(json.ID, 0, "id").Message("id不能小于0")
	for _, err := range valid.Errors {
		c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
		return
	}

	c.JSON(http.StatusOK, util.Success(make(map[string]interface{}), "新闻评论成功"))
}

type CollectNewsJson struct {
	ID int `json:"id" binding:"required"`
}

// CollectNews 收藏新闻
func CollectNews(c *gin.Context) {
	var json CollectNewsJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorFail(e.ERROR_ADD_FAIL, err.Error(), nil))
		return
	}

	valid := validation.Validation{}
	valid.Min(json.ID, 0, "id").Message("id不能小于0")
	for _, err := range valid.Errors {
		c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
		return
	}

	c.JSON(http.StatusOK, util.Success(make(map[string]interface{}), "新闻收藏成功"))
}

type DeleteCommentJson struct {
	ID        int `json:"id" binding:"required"`
	CommentId int `json:"comment_id" binding:"required"`
}

// DeleteComment 删除评论
func DeleteComment(c *gin.Context) {
	var json DeleteCommentJson
	if err := c.ShouldBindJSON(&json); err != nil {
		c.JSON(http.StatusBadRequest, util.ErrorFail(e.ERROR_ADD_FAIL, err.Error(), nil))
		return
	}
	valid := validation.Validation{}
	valid.Min(json.ID, 0, "id").Message("id不能小于0")
	for _, err := range valid.Errors {
		c.JSON(http.StatusOK, util.ErrorFail(e.INVALID_PARAMS, err.Message, nil))
		return
	}

	c.JSON(http.StatusOK, util.Success(make(map[string]interface{}), "评论删除成功"))
}

func Random(c *gin.Context) {
	pType := com.StrTo(c.DefaultQuery("type", "1")).MustInt()
	if pType == 2 {
		if rand.Intn(2) == 0 {
			c.JSON(http.StatusOK, util.Success(1, ""))
		} else {
			c.JSON(http.StatusOK, util.Success(2, ""))
		}
	} else {
		c.JSON(http.StatusOK, util.Success(rand.Intn(2), ""))
	}

}
