package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/gogather/com"
	"strconv"
)

// 登录
type LoginController struct {
	controllers.BaseController
}

func (l *LoginController) Get() {
	l.TplName = "user/login.html"
}
func (l *LoginController) Login() {

	account := l.Input().Get("account")
	password := l.Input().Get("password")
	autoLogin := l.Input().Get("autoLogin") == "on"
	//this.Ctx.WriteString(account + password)

	user, err := models.GetUserByAccount(account)
	//this.Data["username"] = user.Name
	if err != nil {
		beego.Error(err)
	}

	fmt.Print("显示用户名密码显示用户名密码" + user.Name + ";" + user.Password + "输入密码:" + password + "显示用户名密码显示用户名密码")

	if user.Password == com.Md5(password+user.Salt) {
		maxAge := 0
		if autoLogin {
			maxAge = 1<<31 - 1
		}
		sid := strconv.FormatInt(user.Id, 10)
		l.Ctx.SetCookie("userid", sid, maxAge, "/")
		l.Ctx.SetCookie("account", user.Name, maxAge, "/")
		l.Ctx.SetCookie("password", user.Password, maxAge, "/")
		l.Redirect("/", 302)
	}
	return
}

type LogoutController struct {
	controllers.BaseController
}

func (l *LogoutController) Get() {
	l.Ctx.SetCookie("account", "", -1, "/")
	l.Ctx.SetCookie("password", "", -1, "/")
	l.Redirect("/", 302)
}

type RegistorController struct {
	controllers.BaseController
}

func (l *RegistorController) Get() {
	l.TplName = "user/registor.html"
}

func (r *RegistorController) Post() {
	account := r.GetString("account")
	password := r.GetString("password")
	phone := r.GetString("phone")
	_, err := models.AddUser(account, password, phone)
	fmt.Print(err)
	r.Redirect("/", 302)
}
