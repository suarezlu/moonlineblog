package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/suarezlu/moonlineblog/models"
)

type SysController struct {
	beego.Controller
	models.Auth
	Orm orm.Ormer
}

func (this *SysController) Prepare() {
	this.Orm = orm.NewOrm()
	userId := this.GetSession("sysuserid")

	controllerName, actionName := this.Controller.GetControllerAndAction()

	if userId == nil {
		if controllerName != "SysController" || actionName != "Login" {
			this.Redirect("/sys/login", 302)
		}
	} else {
		this.Auth.User.Id = userId.(int)
		err := this.Orm.Read(&this.Auth.User)
		if err == nil {
			this.Auth.Type = "sys"
			this.Auth.IsLogin = true
			this.Data["Auth"] = this.Auth
			this.Data["Controller"] = controllerName
			if actionName != "Home" {
				this.Data["Action"] = actionName
				this.Data["Title"] = "BLOG"
				this.Layout = "layout/sysmain.tpl"
			}
		} else {
			panic(err)
		}
	}
}

// 登录
func (this *SysController) Login() {
	if this.Auth.IsLogin {
		this.Redirect("/sys", 302)
	}
	this.Data["ErrMsg"] = ""

	if this.Ctx.Input.IsPost() {
		username := this.GetString("username", "")
		password := this.GetString("password", "")

		user := models.User{Username: username}
		err := this.Orm.Read(&user, "username")

		if err != nil {
			this.Data["ErrMsg"] = "用户不存在！"
		} else {
			if user.LoginErrTimes >= 3 {
				this.Data["ErrMsg"] = "密码输入错误超过3次"
			} else {
				pwd := Md5(password + user.Salt)
				if pwd == user.Password {
					user.LastIp = this.Ctx.Input.IP()
					user.LoginErrTimes = 0
					this.Orm.Update(&user)

					this.SetSession("sysuserid", user.Id)
					this.Redirect("/sys", 302)
				} else {
					user.LoginErrTimes++
					this.Orm.Update(&user)
					this.Data["ErrMsg"] = "密码错误！"
				}
			}
		}
	}
}

// 退出登录
func (this *SysController) Logout() {
	this.DelSession("sysuserid")
	this.Redirect("/sys/login", 302)
}

// 后台首页
func (this *SysController) Home() {

}

// 分类首页
func (this *SysController) Category() {}

// 分类列表（后台）
func (this *SysController) CategoryList() {
	page, _ := this.GetInt("page", 1)
	limit, _ := this.GetInt("limit", 30)
	offset := (page - 1) * limit
	cnt, _ := this.Orm.QueryTable(new(models.Category)).Count()
	var maps []orm.Params
	this.Orm.QueryTable(new(models.Category)).Limit(limit, offset).Values(&maps)
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "count": cnt, "data": maps}
	this.ServeJSON()
	this.StopRun()
}

// 修改分类（后台）
func (this *SysController) CategoryUpdate() {
	id, _ := this.GetInt("Id")
	name := this.GetString("Name")

	category := models.Category{Id: id}
	err := this.Orm.Read(&category)
	if err == nil {
		category.Name = name
		this.Orm.Update(&category, "Name", "Updated")
	}

	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "data": category}
	this.ServeJSON()
	this.StopRun()
}

// 删除分类（后台）
func (this *SysController) CategoryDel() {
	id, _ := this.GetInt("Id")
	this.Orm.Delete(&models.Category{Id: id})

	this.Data["json"] = map[string]interface{}{"code": 0, "msg": ""}
	this.ServeJSON()
	this.StopRun()
}

// 添加分类
func (this *SysController) CategoryAdd() {
	name := this.GetString("Name")
	if name != "" {
		var category models.Category
		category.Name = name
		_, err := this.Orm.Insert(&category)
		if err == nil {
			this.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "data": category}
		} else {
			this.Data["json"] = map[string]interface{}{"code": 1, "msg": err.Error()}
		}
	} else {
		this.Data["json"] = map[string]interface{}{"code": 1, "msg": "name is empty"}
	}

	this.ServeJSON()
	this.StopRun()
}

// 返回字符串的md5值
func Md5(str string) string {
	if str != "" {
		h := md5.New()
		_, err := io.WriteString(h, str)
		if err == nil {
			return hex.EncodeToString(h.Sum(nil))
		} else {
			panic(err)
		}
	}
	panic(errors.New("str is empty."))
}
