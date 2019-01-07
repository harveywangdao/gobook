package models

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"time"
)

type User struct {
	Id         int       `orm:"auto;column(id)" json:"id" description:"index"`
	Email      string    `orm:"unique;size(100);column(email)" json:"email" description:"email"`
	Token      string    `orm:"size(100);column(token)" json:"token" description:"token"`
	CreateTime time.Time `orm:"type(datetime);column(createTime)" json:"createTime" description:"createTime"`
}

func (u *User) TableName() string {
	return "user_tb"
}

func init() {
	orm.RegisterModel(new(User))
}

func AddUser(user *User) error {
	o := orm.NewOrm()
	user.CreateTime = time.Now()

	id, err := o.Insert(user)
	if err != nil {
		beego.Error(err)
		return err
	}

	_ = id
	return nil
}

func QueryUser(email string) (*User, error) {
	o := orm.NewOrm()
	var user User

	err := o.QueryTable("user_tb").Filter("email", email).One(&user)
	if err != nil {
		beego.Error(err)
		return nil, err
	}

	return &user, nil
}
