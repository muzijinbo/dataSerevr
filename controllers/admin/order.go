package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	"dataServer/utils"
	//"fmt"
	"github.com/astaxie/beego"
)

type OrderController struct {
	controllers.BaseController
}

func (this *OrderController) Get() {
	this.Data["IsLogin"] = utils.CheckAccount(this.Ctx)

	//设置userID
	CK, err := this.Ctx.Request.Cookie("userid")
	if err != nil {
		beego.Error(err)
	}
	userId := CK.Value

	models.CreateOrder(userId)

	//得到账单
	order, err2 := models.GetOrder(userId)
	//fmt.Println("能否有数据？", order.Data)
	this.Data["order"] = order
	if err2 != nil {
		beego.Error(err2)
	}

	this.Data["finished"], err = models.GetFinishedOrder(userId)
	this.Data["unfinished"], err = models.GetUnfinishedOrder(userId)
	this.Data["canceled"], err = models.GetCanceledOrder(userId)
	//获取推荐产品
	this.Data["similarData"], err = models.GetNewData(3)

	this.TplName = "function/order.html"
}

type PayOrderController struct {
	controllers.BaseController
}

func (p *PayOrderController) Get() {
	orderId := p.Input().Get("orderid")
	models.Pay(orderId)
	p.Redirect("/mybills", 302)
}
