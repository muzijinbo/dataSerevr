package models

import (
	"github.com/astaxie/beego/orm"
)

type CartItem struct {
	Id   string `orm:"pk"`
	User *User  `orm:"rel(fk)"`
	Data *Data  `orm:"rel(fk)"`
	Num  int32
}

func init() {
	orm.RegisterModel(new(CartItem))
}

func GetCart(OwnerId string) ([]*CartItem, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("cartItem")
	var cartItems []*CartItem
	_, err := qs.Filter("User", OwnerId).RelatedSel().All(&cartItems)
	return cartItems, err
}

func AddToCart(UserId string, DataId string) (bool, error) {
	o := orm.NewOrm()
	user, err := GetUserById(UserId)
	data := new(Data)
	data, err = GetDataById(DataId)
	str := UserId + "_" + DataId

	cartitem := &CartItem{
		Id:   str,
		User: user,
		Data: data,
		Num:  1,
	}
	_, err = o.Insert(cartitem)
	if err != nil {
		return false, err
	}
	AddActivity(user, data, ADD_CART)
	AddCartNum(DataId)
	return true, err
}

func IsCartItemExist(ID string) bool {
	o := orm.NewOrm()

	qs := o.QueryTable("cartItem")
	IsExist := qs.Filter("Id", ID).Exist()

	return IsExist
}
func DeleteCartItemById(Id string) bool {
	o := orm.NewOrm()
	qs := o.QueryTable("cartItem")
	_, err := qs.Filter("Id", Id).Delete()
	if err != nil {
		return false
	}
	return true
}
