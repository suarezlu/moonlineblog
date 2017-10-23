package controllers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/suarezlu/moonlineblog/models"
)

type BaseController struct {
	beego.Controller
	Orm orm.Ormer
}

func (this *BaseController) Prepare() {
	this.Orm = orm.NewOrm()
	if !this.Ctx.Input.IsAjax() {
		this.Layout = "layout/main.tpl"

		this.LayoutSections = make(map[string]string)

		var categories []orm.Params
		i, _ := this.Orm.QueryTable(new(models.Category)).Values(&categories)
		this.Data["Categories"] = categories
		this.Data["I"] = i
		this.LayoutSections["Nav"] = "common/nav.tpl"

		this.Data["Title"] = "BLOG"
	}

	//	this.LayoutSections = make(map[string]string)
	//	this.LayoutSections["Navbar"] = "common/navbar.html"
	//	this.LayoutSections["Sidebar"] = "common/sidebar.html"

	//	if this.GetSession("username") != nil {
	//		this.Data["Islogin"] = true
	//		this.Data["Username"] = this.GetSession("username")
	//		this.User.Id = 1
	//	} else {
	//		this.Data["Islogin"] = false
	//		this.Data["Username"] = ""
	//	}

	//	this.Data["Title"] = "suarez blog"
	//	this.Data["CatList"], _, this.Data["num"] = CategortList()
}
