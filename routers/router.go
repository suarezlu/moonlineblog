package routers

import (
	"github.com/suarezlu/moonlineblog/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
}
