package models

import (
	"github.com/astaxie/beego/orm"
	"time"
)

type Activity struct {
	Id   int64
	User *User `orm:"rel(fk)"`
	Data *Data `orm:"rel(fk)"`
	Op   int32 //1 浏览数据 2.加入购物车 3.收藏 4.购买数据
	Time time.Time
}

const (
	BROWSE_DATA    int32 = 1 //1 浏览数据 2.加入购物车 3.收藏 4.购买数据
	ADD_CART       int32 = 2
	ADD_COLLECTION int32 = 3
	BUY_DATA       int32 = 4
)

func init() {
	orm.RegisterModel(new(Activity))
}

func AddActivity(user *User, data *Data, op int32) (int64, error) {
	o := orm.NewOrm()
	activity := new(Activity)
	activity.User = user
	activity.Data = data
	activity.Op = op
	activity.Time = time.Now()

	return o.Insert(activity)
}

func GetBrowseDatas(OwnerId string, num int) ([]*Data, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("activity")
	activities := make([]*Activity, 0)
	datas := make([]*Data, 0)
	_, err := qs.Filter("User", OwnerId).Filter("Op", BROWSE_DATA).RelatedSel().Limit(num).All(&activities)
	for _, v := range activities {
		datas = append(datas, v.Data)
	}
	return datas, err
}
