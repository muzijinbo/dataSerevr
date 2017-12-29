package models

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	"time"
)

type Data struct {
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
	Author      *User    `orm:"rel(fk)"`
	Order       []*Order `orm:"reverse(many)"`
	BrowseNum   int32    //浏览量
	CollectNum  int32    //收藏量
	CartNum     int32    //加入购物车次数
	PayNum      int32    //销售量
}

func init() {
	orm.RegisterModel(new(Data))
}

func AddData(classId, name, primary, introduce, path, autherId string) (int64, error) {
	o := orm.NewOrm()
	data := new(Data)
	data.Name = name
	data.Primary = primary
	data.Introduce = introduce
	data.Location = path
	data.DType = classId
	data.Created = time.Now()
	data.Author, _ = GetUserById(autherId)
	a, b := o.Insert(data)
	fmt.Println(a, b)
	return a, b
}

func GetDataById(id string) (*Data, error) {
	o := orm.NewOrm()
	//users := make([]*User, 0)

	qs := o.QueryTable("data")
	data := new(Data)
	err := qs.Filter("id", id).One(data)
	//_, err := qs.All(&users)
	return data, err
}

//获得热销商品
func GetHotData(num int) ([]*Data, error) {
	o := orm.NewOrm()
	datas := make([]*Data, 0)
	qs := o.QueryTable("data")
	qs.OrderBy("-PayNum")
	qs = qs.Limit(num)
	_, err := qs.All(&datas)
	return datas, err
}

//获得新上架商品
func GetNewData(num int) ([]*Data, error) {
	o := orm.NewOrm()
	datas := make([]*Data, 0)
	qs := o.QueryTable("data")
	qs.OrderBy("-Id")
	qs = qs.Limit(num)
	_, err := qs.All(&datas)
	return datas, err
}

//获得所有商品
func GetAllData() ([]*Data, error) {
	o := orm.NewOrm()
	datas := make([]*Data, 0)
	qs := o.QueryTable("data")
	_, err := qs.All(&datas)
	return datas, err
}

//数据浏览量加一
func AddBrowseNum(dataid string) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("Data").Filter("id", dataid).Update(orm.Params{
		"BrowseNum": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}

//数据收藏量加一
func AddCollectNum(dataid string) error {

	o := orm.NewOrm()
	_, err := o.QueryTable("Data").Filter("id", dataid).Update(orm.Params{
		"CollectNum": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}

//购物车添加次数加一
func AddCartNum(dataid string) error {

	o := orm.NewOrm()
	_, err := o.QueryTable("Data").Filter("id", dataid).Update(orm.Params{
		"CartNum": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}

//销售量加一
func AddPayNum(dataid int64) error {
	o := orm.NewOrm()
	_, err := o.QueryTable("Data").Filter("id", dataid).Update(orm.Params{
		"PayNum": orm.ColValue(orm.ColAdd, 1),
	})
	return err
}

func GetDataList(page, pageSize int, filters ...interface{}) ([]*Data, int64) {
	offset := (page - 1) * pageSize

	Datas := make([]*Data, 0)

	query := orm.NewOrm().QueryTable("Data")
	if len(filters) > 0 {
		l := len(filters)
		for k := 0; k < l; k += 2 {
			query = query.Filter(filters[k].(string), filters[k+1])
		}
	}
	total, _ := query.Count()
	query.OrderBy("Id").Limit(pageSize, offset).All(&Datas)

	return Datas, total
}
