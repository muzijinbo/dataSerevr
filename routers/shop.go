package routers

import (
	"dataServer/controllers/admin"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/shop", &admin.ShopController{})

}
