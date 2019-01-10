package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type Article struct {
	Id         int64     `orm:"auto;column(id)" json:"id" description:"index"`
	UserId     int64     `orm:"column(user_id)" json:"userId" description:"userId"`
	Title      string    `orm:"size(100);column(title)" json:"title" description:"title"`
	Text       string    `orm:"size(10000);column(text)" json:"text" description:"text"`
	Tags       string    `orm:"size(100);column(tags)" json:"tags" description:"tags"`
	CreateTime time.Time `orm:"type(datetime);column(createTime)" json:"createTime" description:"createTime"`
}

func (a *Article) TableName() string {
	return "article_tb"
}

func init() {
	orm.RegisterModel(new(Article))
}

func AddArticle(a *Article) error {
	o := orm.NewOrm()
	a.CreateTime = time.Now()

	id, err := o.Insert(a)
	if err != nil {
		beego.Error(err)
		return err
	}

	_ = id
	return nil
}

func QueryAllArticles() ([]*Article, error) {
	o := orm.NewOrm()
	var articles []*Article

	_, err := o.QueryTable("article_tb").OrderBy("-id").All(&articles)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return articles, nil
}

func QueryUserArticles(userId int64) ([]*Article, error) {
	o := orm.NewOrm()
	var articles []*Article

	_, err := o.QueryTable("article_tb").Filter("user_id", userId).OrderBy("-id").All(&articles)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return articles, nil
}

func QueryArticle(articleId int64) (*Article, error) {
	o := orm.NewOrm()
	var article Article

	err := o.QueryTable("article_tb").Filter("id", articleId).One(&article)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return &article, nil
}
