package routers

import (
	"dataServer/controllers"
	"dataServer/controllers/admin"
	"dataServer/controllers/index"
	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &index.MainController{})
	beego.Router("/myaccount", &index.MyaccountController{})
	beego.Router("/uploaddata", &admin.UploadController{})
	beego.Router("/uploaddata/loadclass", &admin.UploadController{}, "post:GetClasses")
	beego.Router("/uploaddata/upload", &admin.UploadController{}, "post:Upload")
	beego.Router("/crawl", &admin.CrawlController{})
	beego.Router("/crawldata", &admin.CrawlController{}, "post:Crawl")

	beego.Router("/mybills", &admin.BillsController{})

	beego.Router("/index/details", &admin.CheckDetailsController{})
	beego.Router("/cart", &admin.CartController{})
	beego.Router("/order", &admin.OrderController{})
	beego.Router("/order/pay", &admin.PayOrderController{})

	beego.Router("/test", &controllers.TestController{})
	beego.Router("/download", &admin.DownloadController{})

}
