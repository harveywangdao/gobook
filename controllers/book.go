package controllers

import (
	"github.com/astaxie/beego"
	"gobook/models"
	"strconv"
)

type ArticleController struct {
	beego.Controller
}

// @Title AddArticle
// @Description add article
// @Param   title   formData    string  true    "title"
// @Param   text    formData    string  true    "text"
// @Param   tags    formData    string  true    "tags"
// @Success 200 {object} models.ApiResult
// @Failure 403 body is empty
// @router /addArticle [post]
func (u *ArticleController) AddArticle() {
	var ret models.ApiResult
	ret.Status = false

	title := u.GetString("title")
	text := u.GetString("text")
	tags := u.GetString("tags")

	loginSession := u.GetSession(UserLoginSession)
	if loginSession == nil {
		ret.Msg = "user unlogin"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	userId, ok := loginSession.(int64)
	if !ok {
		ret.Msg = "get userId failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	article := models.Article{
		UserId: userId,
		Title:  title,
		Text:   text,
		Tags:   tags,
	}

	err := models.AddArticle(&article)
	if err != nil {
		ret.Msg = "add article failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	ret.Status = true
	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title QueryAllArticles
// @Description query all articles
// @Success 200 {object} models.ArticlesResult
// @Failure 403 no data
// @router /queryAllArticles [get]
func (u *ArticleController) QueryAllArticles() {
	var ret models.ArticlesResult
	ret.Status = false

	articles, err := models.QueryAllArticles()
	if err != nil {
		ret.Msg = "query all articles failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	ret.Status = true
	ret.Data = articles

	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title QueryUserArticles
// @Description query user articles
// @Success 200 {object} models.ArticlesResult
// @Failure 403 no data
// @router /queryUserArticles [get]
func (u *ArticleController) QueryUserArticles() {
	var ret models.ArticlesResult
	ret.Status = false

	loginSession := u.GetSession(UserLoginSession)
	if loginSession == nil {
		ret.Msg = "user unlogin"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	userId, ok := loginSession.(int64)
	if !ok {
		ret.Msg = "get userId failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	articles, err := models.QueryUserArticles(userId)
	if err != nil {
		ret.Msg = "query user articles failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	ret.Status = true
	ret.Data = articles

	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title QueryArticle
// @Description query one article
// @Param   articleId    path    string  true    "articleId"
// @Success 200 {object} models.ArticleResult
// @Failure 403 no data
// @router /queryOneArticle/:articleId [get]
func (u *ArticleController) QueryArticle() {
	var ret models.ArticleResult
	ret.Status = false

	articleIdStr := u.Ctx.Input.Param(":articleId")
	articleId, err := strconv.ParseInt(articleIdStr, 10, 64)
	if err != nil {
		ret.Msg = "articleId param error"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	article, err := models.QueryArticle(articleId)
	if err != nil {
		ret.Msg = "query article failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	ret.Status = true
	ret.Data = article

	u.Data["json"] = ret
	u.ServeJSON()
}
