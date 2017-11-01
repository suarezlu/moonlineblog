package controllers

import (
	"github.com/suarezlu/moonlineblog/models"
)

type BlogController struct {
	BaseController
}

func (this *BlogController) Index() {
	categoryId, _ := this.GetInt("cat", 0)
	page, _ := this.GetInt("page", 1)
	limit, _ := this.GetInt("limit", 20)

	offset := (page - 1) * limit
	qs := this.Orm.QueryTable(new(models.Article)).OrderBy("-release_time")
	if categoryId > 0 {
		qs = qs.Filter("category_id", categoryId)
	}
	cnt, _ := qs.Count()
	var list []models.Article
	qs.RelatedSel().Limit(limit, offset).All(&list)
	this.Data["List"] = list
	this.Data["PageInfo"] = map[string]interface{}{"CategoryId": categoryId, "Page": page, "Limit": limit, "Count": cnt}
}
