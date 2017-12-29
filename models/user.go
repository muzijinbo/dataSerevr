package models

import (
	"github.com/astaxie/beego/orm"
	"github.com/gogather/com"
	"time"
)

type User struct {
	Id       int64
	Name     string
	Phone    string
	Password string
	Created  time.Time `orm:"index;auto_now_add;type(datetime)"`
	Level    int32     `orm:"index"`
	UType    int32
	Salt     string
	Balance  int64
}

func init() {
	orm.RegisterModel(new(User))
}
func GetUserById(id string) (*User, error) {
	o := orm.NewOrm()
	//users := make([]*User, 0)
	qs := o.QueryTable("user")
	user := new(User)
	err := qs.Filter("id", id).One(user)
	//_, err := qs.All(&users)
	return user, err
}

func GetUserByAccount(account string) (*User, error) {
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	user := new(User)
	err := qs.Filter("name", account).One(user)
	//_, err := qs.All(&users)
	return user, err
}

func AddUser(account, password, phone string) (int64, error) {
	o := orm.NewOrm()
	user := new(User)
	user.Name = account
	user.Created = time.Now()
	user.Salt = com.RandString(10)
	user.Phone = phone
	user.Password = com.Md5(password + user.Salt)
	return o.Insert(user)
}

func MinusBalance(num float32, userid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("User").Filter("Id", userid).Update(orm.Params{
		"Balance": orm.ColValue(orm.ColMinus, num),
	})
	return err
}
