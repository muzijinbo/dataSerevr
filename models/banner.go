package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type BannerData struct {
	Id          int64
	Location    string
	ImgLocation string
	Name        string
	Key1        string
	Key2        string
	Key3        string
	Primary     string
	Number      int
	Created     time.Time `orm:"index"` //创建时间
	Size        float32   //大小
	Introduce   string    //简介
	Attachment  string    `orm:"size(5000)"`
	DType       string    //类型
	Price       float32   //原价
	PriceNow    float32   //现价
	Popular     int       //火热程度
	Tags        string
	Author      *User `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(BannerData))
}

func GetBannerData() ([]*BannerData, error) {
	o := orm.NewOrm()
	datas := make([]*BannerData, 0)
	qs := o.QueryTable("banner_data")
	qs = qs.Limit(4)
	_, err := qs.All(&datas)
	return datas, err
}
