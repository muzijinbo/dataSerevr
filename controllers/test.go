package controllers

import ()

// 登录
type TestController struct {
	BaseController
}

func (t *TestController) Get() {
	t.TplName = "test.html"
}

func (t *TestController) Post() {
	t.Data["json"] = map[string]interface{}{"result": true, "msg": "成功", "refer": "/myaccount"}
	t.ServeJSON()
}

// 登录
type PayController struct {
	BaseController
}

func (t *PayController) Get() {
	t.TplName = "test.html"
}

func (t *PayController) Post() {

	/*alipay := alipay.Client{
		Partner:   "1234567",           // 合作者ID
		Key:       "1234567",           // 合作者私钥
		ReturnUrl: "/",                 // 同步返回地址
		NotifyUrl: "/",                 // 网站异步返回地址
		Email:     "1355974104@qq.com", // 网站卖家邮箱地址
	}
	form := alipay.Form(alipay.Options{
		OrderId:  "123",   // 唯一订单号
		Fee:      99.8,    // 价格
		NickName: "翱翔大空",  // 用户昵称，支付页面显示用
		Subject:  "充值100", // 支付描述，支付页面显示用
	})*/

}
