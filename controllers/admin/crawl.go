package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	//"dataServer/utils"
	//"fmt"
	"github.com/astaxie/beego"
	"webcrawler/demo"
)

type CrawlController struct {
	controllers.BaseController
}

func (c *CrawlController) Get() {
	classes, err := models.GetClassifyAll(1)
	if err != nil {
		beego.Error(err)
	}
	c.Data["classes"] = classes
	c.TplName = "myaccount/crawldata.html"
}

func (c *CrawlController) Crawl() {
	//classes2 := u.Input().Get("classes2")
	//name := u.GetString("name")
	//primary := u.GetString("primary")
	//introduce := u.GetString("introduce")
	//filepath := beego.AppConfig.String("filepath")
	address := c.GetString("address")
	demo.Crewl(address)
	c.Data["json"] = map[string]interface{}{"result": false, "msg": "获取文件失败", "refer": "/uploaddata"}
	c.ServeJSON()
}
