package routers

import (
	"github.com/astaxie/beego"
	"github.com/suarezlu/moonlineblog/controllers"
)

func init() {
	beego.Router("/", &controllers.BlogController{}, "get:Index")
	beego.Router("/blog/test", &controllers.BlogController{}, "get:Test")

	// 管理后台
	beego.Router("/sys/login", &controllers.SysController{}, "*:Login")
	beego.Router("/sys/logout", &controllers.SysController{}, "get:Logout")
	beego.Router("/sys", &controllers.SysController{}, "get:Home")
	// 分类
	beego.Router("/sys/category", &controllers.SysController{}, "get:Category")
	beego.Router("/sys/categorylist", &controllers.SysController{}, "get:CategoryList")
	beego.Router("/sys/categoryupdate", &controllers.SysController{}, "post:CategoryUpdate")
	beego.Router("/sys/categorydel", &controllers.SysController{}, "post:CategoryDel")
}
