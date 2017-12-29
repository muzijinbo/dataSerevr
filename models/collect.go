package models

import (
	"github.com/astaxie/beego/orm"
)

type Collect struct {
	Id   int
	User *User `orm:"rel(fk)"`
	Data *Data `orm:"rel(fk)"`
}

func init() {
	orm.RegisterModel(new(Collect))
}

func AddCollext(UserId string, DataId string) error {
	o := orm.NewOrm()
	user, err := GetUserById(UserId)
	data := new(Data)
	data, err = GetDataById(DataId)

	cartitem := &CartItem{
		User: user,
		Data: data,
	}
	_, err = o.Insert(cartitem)
	_, err = AddActivity(user, data, ADD_COLLECTION)
	err = AddCollectNum(DataId)
	return err
}
