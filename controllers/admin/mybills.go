package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	"github.com/astaxie/beego"
)

type BillsController struct {
	controllers.BaseController
}

func (b *BillsController) Get() {
	//获取userID
	CK, err := b.Ctx.Request.Cookie("userid")
	if err != nil {
		beego.Error(err)
	}
	userId := CK.Value
	b.Data["finished"], err = models.GetFinishedOrder(userId)
	b.Data["unfinished"], err = models.GetUnfinishedOrder(userId)
	b.Data["canceled"], err = models.GetCanceledOrder(userId)
	b.TplName = "myaccount/mybills.html"
}
