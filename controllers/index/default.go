package index

import (
	"dataServer/controllers"
	"dataServer/models"
	"dataServer/utils"
	"github.com/astaxie/beego"
	//"fmt"
)

type MainController struct {
	controllers.BaseController
}

func (c *MainController) Get() {
	c.Data["IsLogin"] = utils.CheckAccount(c.Ctx)
	var err error
	c.Data["bannerdata"], err = models.GetBannerData()
	//c.Data["lastData"], err = models.GetNewData()
	c.Data["hotdata"], err = models.GetHotData(3)
	c.Data["newdata8"], err = models.GetNewData(8)
	c.Data["newdata"], err = models.GetNewData(3)
	//设置userID
	CK, err := c.Ctx.Request.Cookie("userid")
	if err != nil {
		beego.Error(err)
	}
	userId := CK.Value
	c.Data["browsedata"], err = models.GetBrowseDatas(userId, 3)

	c.TplName = "index.html"
}

type MyaccountController struct {
	controllers.BaseController
}

func (c *MyaccountController) Get() {
	c.Data["IsLogin"] = utils.CheckAccount(c.Ctx)

	c.TplName = "myaccount/uploaddata.html"
}
