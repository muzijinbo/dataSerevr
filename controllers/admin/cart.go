package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	"dataServer/utils"
	"fmt"
	"github.com/astaxie/beego"
)

type CartController struct {
	controllers.BaseController
}

func (this *CartController) Get() {
	this.Data["IsLogin"] = utils.CheckAccount(this.Ctx)

	//设置userID
	CK, err := this.Ctx.Request.Cookie("userid")
	if err != nil {
		beego.Error(err)
	}
	userId := CK.Value
	this.Data["userId"] = CK.Value
	fmt.Println("显示CKValue:" + CK.Value)

	//得到购物车明细
	items, err2 := models.GetCart(userId)
	this.Data["items"] = items
	for _, v := range items {
		fmt.Println("显示条目" + v.Data.ImgLocation)
	}
	if err2 != nil {
		beego.Error(err2)
	}

	//获取推荐产品
	this.Data["similarData"], err = models.GetNewData(3)

	//一些操作
	op := this.Input().Get("op")
	dataId := this.Input().Get("dataId")
	var IsSuccess bool
	switch op {
	case "add":
		IsExist := models.IsCartItemExist(userId + "_" + dataId)
		if !IsExist {
			IsSuccess, err = models.AddToCart(userId, dataId)
			if !IsSuccess {
				fmt.Println("插入失败")
			}
		}
		if IsExist {
			fmt.Println("已经添加该商品至购物车了")
		}
		this.Redirect("/", 302)
		return
	case "del":
		dataItemId := this.Input().Get("dataItemId")
		IsSuccess = models.DeleteCartItemById(dataItemId)
		if IsSuccess {
			this.Redirect("/cart", 302)
		}
	}
	this.TplName = "function/cart.html"
}
