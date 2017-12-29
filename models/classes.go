package models

import (
	//"fmt"
	"github.com/astaxie/beego/orm"
)

type Classify struct {
	Id       int
	Name     string
	Father   *Classify   `orm:"rel(fk)"`
	Children []*Classify `orm:"reverse(many)"`
}

func init() {
	orm.RegisterModel(new(Classify))
}

func GetClassifyAll(fatherId int) ([]*Classify, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("classify")
	var classifies []*Classify
	_, err := qs.Filter("Father", fatherId).All(&classifies)

	return classifies, err
}
