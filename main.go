package main

import (
	_ "dataServer/initial"
	_ "dataServer/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

func main() {
	orm.RunSyncdb("default", false, true)
	beego.Run()
}
