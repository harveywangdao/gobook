package controllers

import (
	"gobook/models"

	"github.com/astaxie/beego"
)

const (
	UserLoginSession = "UserLoginSession"
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

// @Title Login
// @Description login user into the system
// @Param	email	formData	string	true	"email"
// @Param	token	formData	string	true	"password token"
// @Success 200 {object} models.ApiResult
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	var ret models.ApiResult
	ret.Status = false

	email := u.GetString("email")
	token := u.GetString("token")

	user, err := models.QueryUser(email)
	if err != nil {
		ret.Msg = "user not existed"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	loginSession := u.GetSession(UserLoginSession)
	if loginSession != nil {
		ret.Msg = "user already logined"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	if user.Token != token {
		ret.Msg = "password error"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	u.SetSession(UserLoginSession, user.Id)

	ret.Status = true
	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title logout
// @Description user logout
// @Success 200 {object} models.ApiResult
// @Failure 403 user not exist
// @router /logout [post]
func (u *UserController) Logout() {
	var ret models.ApiResult
	ret.Status = false

	loginSession := u.GetSession(UserLoginSession)
	if loginSession == nil {
		ret.Msg = "user unlogin"
		u.Data["json"] = ret
		u.ServeJSON()
		return
	}

	u.DelSession(UserLoginSession)

	ret.Status = true
	u.Data["json"] = ret
	u.ServeJSON()
}

// @Title isLogined
// @Description query login status
// @Success 200 {object} models.ApiResult
// @Failure 403 user not exist
// @router /isLogined [get]
func (u *UserController) IsLogined() {
	var ret models.ApiResult

	loginSession := u.GetSession(UserLoginSession)
	if loginSession == nil {
		ret.Status = false
	} else {
		ret.Status = true
	}

	u.Data["json"] = ret
	u.ServeJSON()
}
