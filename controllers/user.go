package controllers

import (
	"github.com/astaxie/beego"
	"gobook/models"
)

type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	email	formData	string	true	"email"
// @Param	token	formData	string	true	"password token"
// @Success 200 {object} models.ApiResult
// @Failure 403 body is empty
// @router /register [post]
func (u *UserController) Register() {
	var ret models.ApiResult
	ret.Status = false

	email := u.GetString("email")
	token := u.GetString("token")
	user := models.User{
		Email: email,
		Token: token,
	}

	err := models.AddUser(&user)
	if err != nil {
		ret.Msg = "insert user failure"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	ret.Status = true
	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /getAllUsers [get]
func (u *UserController) GetAll() {
	u.Data["json"] = "users"
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	u.Data["json"] = "user"
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	u.Data["json"] = "uu"
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	u.Data["json"] = "delete success!"
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {
	u.Data["json"] = "user not exist"
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
