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
	beego.Router("/sys/pwd", &controllers.SysController{}, "*:Pwd")
	beego.Router("/sys/upload", &controllers.SysController{}, "*:Upload")
	// 分类
	beego.Router("/sys/category", &controllers.SysController{}, "get:Category")
	beego.Router("/sys/categorylist", &controllers.SysController{}, "get:CategoryList")
	beego.Router("/sys/categoryupdate", &controllers.SysController{}, "post:CategoryUpdate")
	beego.Router("/sys/categorydel", &controllers.SysController{}, "post:CategoryDel")
	beego.Router("/sys/categoryadd", &controllers.SysController{}, "post:CategoryAdd")
	// 文章
	beego.Router("/sys/articles", &controllers.SysController{}, "get:Articles")
	beego.Router("/sys/articlelist", &controllers.SysController{}, "get:ArticleList")
	beego.Router("/sys/article/:id([0-9]+)", &controllers.SysController{}, "get:Article")
	beego.Router("/sys/articlesave", &controllers.SysController{}, "post:ArticleSave")
	beego.Router("/sys/articledel", &controllers.SysController{}, "post:ArticleDel")
}
