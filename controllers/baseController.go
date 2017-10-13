package controllers

import (
	"github.com/astaxie/beego"
	. "github.com/suarezlu/moonlineblog/models"
)

type BaseController struct {
	beego.Controller
	Auth
}

func (this *BaseController) Prepare() {
	this.Layout = "layout/main.tpl"

	this.Data["Title"] = "BLOG"

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
