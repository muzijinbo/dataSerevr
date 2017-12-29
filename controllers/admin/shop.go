package admin

import (
	"dataServer/controllers"
	"dataServer/models"
	"github.com/astaxie/beego"
)

type ShopController struct {
	controllers.BaseController
}

func (s *ShopController) Get() {

	var classfies []*models.Classify
	var err error
	classfies, err = models.GetClassifyAll(1)
	if err != nil {
		beego.Error(err)
	}
	for classItem := range classfies {
		classfies[classItem].Children, err = models.GetClassifyAll(classfies[classItem].Id)
	}
	s.Data["classfies"] = classfies
	s.Data["alldata"], err = models.GetAllData()

	page, _ := s.GetInt("page")
	if page < 1 {
		page = 1
	}
	filters := make([]interface{}, 0)
	datas, counts := models.GetDataList(page, 4, filters...)
	s.Data["pagedatas"] = datas
	s.Data["counts"] = counts
	s.TplName = "shop.html"
}
