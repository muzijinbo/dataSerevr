package initial

import (
	//"dataServer/utils"
	_ "dataServer/models"
	"github.com/astaxie/beego"
	"github.com/beego/i18n"
)

func InitTplFunc() {
	beego.AddFuncMap("i18n", i18n.Tr)
}
