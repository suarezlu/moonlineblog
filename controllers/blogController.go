package controllers

type BlogController struct {
	BaseController
}

func (this *BlogController) Index() {
	this.TplName = "blog/index.tpl"
}

func (this *BlogController) Test() {
	this.TplName = "blog/test.tpl"
}
