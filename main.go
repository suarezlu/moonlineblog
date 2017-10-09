package main

import (
	"github.com/astaxie/beego"
	_ "github.com/suarezlu/moonlineblog/models"
	_ "github.com/suarezlu/moonlineblog/routers"
)

func main() {
	//models.Init()
	beego.BConfig.WebConfig.TemplateLeft = "{{{"
	beego.BConfig.WebConfig.TemplateRight = "}}}"
	beego.Run()
}
