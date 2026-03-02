package controllers

import (
	"encoding/json"
	"frontend-backend/models"

	"github.com/astaxie/beego"
)

type ArticleController struct {
	beego.Controller
}

func (c *ArticleController) CreateArticle() {
	// 解析请求体
	var req models.Article
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.jsonError(400, "无效的请求数据: "+err.Error())
		return
	}

	// 验证请求数据
	if req.Title == "" || req.Content == "" || req.Author == "" {
		c.jsonError(400, "标题、内容和作者不能为空")
		return
	}

	// 创建新文章
	_, err := models.AddArticle(req)
	if err != nil {
		c.jsonError(500, "创建文章失败: "+err.Error())
		return
	}

	c.jsonSuccess(200, "文章创建成功", nil)
}

func (c *ArticleController) GetArticleList() {
	var req struct {
		Page int `json:"page"`
		Size int `json:"size"`
	}
	req.Page = 1
	req.Size = 10
	if err := json.NewDecoder(c.Ctx.Request.Body).Decode(&req); err != nil {
		c.jsonError(400, "无效的请求数据: "+err.Error())
		return
	}
	var articleList []models.Article
	articleList, _ = models.QueryArticleWithPage(req.Page, req.Size)
	c.jsonSuccess(200, "文章列表查询成功", articleList)
}

func (c *ArticleController) jsonError(i int, s string) {
	c.Ctx.Output.SetStatus(i)
	c.Data["json"] = map[string]interface{}{
		"code": i,
		"msg":  s,
	}
	c.ServeJSON()
}

func (c *ArticleController) jsonSuccess(statusCode int, message string, data interface{}) {
	c.Data["json"] = map[string]interface{}{
		"code":    1,
		"success": true,
		"message": message,
		"data":    data,
	}
	c.Ctx.Output.SetStatus(statusCode)
	c.ServeJSON()
}
