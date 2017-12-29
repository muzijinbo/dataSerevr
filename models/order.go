package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Order struct {
	Id      int64
	User    *User   `orm:"rel(fk)"`
	Data    []*Data `orm:"rel(m2m)"`
	State   int32
	Created time.Time `orm:"index;auto_now_add;type(datetime)"`
	Price   float32
}

func init() {
	orm.RegisterModel(new(Order))
}

func CreateOrder(UserId string) (int64, error) {
	items, err := GetCart(UserId)
	var datas []*Data

	if len(items) > 0 {
		var price float32
		price = 0
		for _, v := range items {
			DeleteCartItemById(v.Id)
			datas = append(datas, v.Data)
			price = price + v.Data.PriceNow
		}
		fmt.Println("我的数据", datas)

		o := orm.NewOrm()
		Order := new(Order)
		Order.Data = datas
		Order.State = 1
		user, err := GetUserById(UserId)
		Order.User = user
		Order.Price = price
		Order.Created = time.Now()
		//fmt.Println("我的订单", Order)
		index, err := o.Insert(Order)
		for _, v := range datas {
			v.Order = append(v.Order, Order)
		}
		num, err := o.QueryM2M(Order, "data").Add(Order.Data)
		fmt.Println("添加多对多关系", num, err)
		return index, err
	} else {
		return 0, err
	}
}

func GetOrder(OwnerId string) (*Order, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Order")
	order := new(Order)
	err := qs.Filter("User", OwnerId).Filter("State", 1).RelatedSel().One(order)
	//err = o.Read(&order)
	_, err = orm.NewOrm().LoadRelated(order, "data")
	return order, err
}

func ChangeOrderState(orderid int64, state int32) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("Order").Filter("id", orderid).Update(orm.Params{
		"State": state,
	})
	return err
}

//未支付订单
func GetUnfinishedOrder(OwnerId string) ([]*Order, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Order")
	orders := make([]*Order, 0)
	_, err := qs.Filter("User", OwnerId).Filter("State", 1).RelatedSel().All(&orders)
	for _, v := range orders {
		_, err = orm.NewOrm().LoadRelated(v, "data")
	}
	return orders, err
}

//已完成订单
func GetFinishedOrder(OwnerId string) ([]*Order, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Order")
	orders := make([]*Order, 0)
	_, err := qs.Filter("User", OwnerId).Filter("State", 2).RelatedSel().All(&orders)
	for _, v := range orders {
		_, err = orm.NewOrm().LoadRelated(v, "data")
	}
	return orders, err
}

//已取消订单
func GetCanceledOrder(OwnerId string) ([]*Order, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("Order")
	orders := make([]*Order, 0)
	_, err := qs.Filter("User", OwnerId).Filter("State", 3).RelatedSel().All(&orders)
	for _, v := range orders {
		_, err = orm.NewOrm().LoadRelated(v, "data")
	}
	return orders, err
}

func Pay(orderId string) error {
	//var err error
	o := orm.NewOrm()
	qs := o.QueryTable("Order")
	order := new(Order)
	err := qs.Filter("Id", orderId).RelatedSel().One(order)
	_, err = orm.NewOrm().LoadRelated(order, "data")
	for _, v := range order.Data {
		AddPayNum(v.Id)
		AddActivity(order.User, v, BUY_DATA)
	}
	MinusBalance(order.Price, order.User.Id)
	ChangeOrderState(order.Id, 2)
	return err
}
