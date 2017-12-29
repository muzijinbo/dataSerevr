package routers

import (
	"dataServer/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/login", &admin.LoginController{})
	beego.Router("/login/post", &admin.LoginController{}, "post:Login")
	beego.Router("/logout", &admin.LogoutController{})
	beego.Router("/registor", &admin.RegistorController{})
}
