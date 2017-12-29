package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	"dataServer/utils"
	"fmt"
	"github.com/astaxie/beego"
	"strconv"
)

// 分类
type UploadController struct {
	controllers.BaseController
}

func (u *UploadController) Get() {
	classes, err := models.GetClassifyAll(1)
	fmt.Println(classes)
	fmt.Println(err)

	/*for classItem := range classes {
		classes[classItem].Children, err = models.GetClassifyAll(classes[classItem].Id)
		fmt.Println("显示分类：")
		fmt.Println(classItem)
	}*/
	u.Data["classes"] = classes
	u.TplName = "myaccount/uploaddata.html"
	//u.TplName = "myaccount/uploaddata.html"
}
func (u *UploadController) GetClasses() {
	sid := u.GetString("id")
	iid, err := strconv.Atoi(sid)
	classes, err := models.GetClassifyAll(iid)
	fmt.Println(err)
	u.Data["json"] = classes
	u.ServeJSON()
}
func (u *UploadController) Upload() {
	//classes2 := u.GetString("classes2")
	classes2 := u.Input().Get("classes2")
	name := u.GetString("name")
	primary := u.GetString("primary")
	introduce := u.GetString("introduce")
	filepath := beego.AppConfig.String("filepath")
	fmt.Println(classes2, name, primary, introduce, filepath)
	f, h, err := u.GetFile("myfile")
	if err != nil {
		u.Data["json"] = map[string]interface{}{"result": false, "msg": "获取文件失败", "refer": "/uploaddata"}
		u.ServeJSON()
	}
	//u.Data["json"] = map[string]interface{}{"result": true, "msg": "获取文件成功", "refer": "/uploaddata"}
	path := filepath + "/" + h.Filename

	f.Close() //关闭上传的文件，不然的话会出现临时文件不能清除的情况
	fmt.Println(path)
	ck, err := u.Ctx.Request.Cookie("userid")
	u.SaveToFile("myfile", path)
	_, err = models.AddData(classes2, name, primary, introduce, path, ck.Value)
	if err != nil {
		u.Data["json"] = map[string]interface{}{"result": false, "msg": "上传数据失败", "refer": "/myaccount"}
	} else {
		u.Data["json"] = map[string]interface{}{"result": true, "msg": "上传数据成功", "refer": "/myaccount"}
	}
	u.ServeJSON()
}

//商品详情
type CheckDetailsController struct {
	controllers.BaseController
}

func (this *CheckDetailsController) Get() {
	dataId := this.Input().Get("id")
	fmt.Println("显示商品id:" + dataId)

	data, err := models.GetDataById(dataId)
	this.Data["singleData"] = data
	fmt.Println("显示数据" + data.Name)
	this.Data["similarData"], err = models.GetNewData(3)
	//this.Data["singleID"] = dataId
	if err != nil {
		beego.Error(err)
	}
	models.AddBrowseNum(dataId)

	CK, err := this.Ctx.Request.Cookie("userid")
	if err != nil {
		beego.Error(err)
	}
	userId := CK.Value
	user, err := models.GetUserById(userId)
	models.AddActivity(user, data, models.BROWSE_DATA)
	this.Data["IsLogin"] = utils.CheckAccount(this.Ctx)
	this.TplName = "function/checkdetiles.html"
}

//下载数据
type DownloadController struct {
	controllers.BaseController
}

func (d *DownloadController) Get() {
	d.Ctx.Output.Download("E:/dataserverfile/LICENSE", "DStore")
}
