package controllers

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"io"
	"math/rand"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

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

// 修改密码
func (this *SysController) Pwd() {
	if this.Ctx.Input.IsPost() {
		pwd := this.GetString("pwd", "")
		newPwd := this.GetString("newpwd")
		repeatPwd := this.GetString("repeatpwd")
		code := 0
		msg := ""
		if this.Auth.Password == Md5(pwd+this.User.Salt) {
			if newPwd == repeatPwd {
				salt := GetRandomString(5)
				msg = salt
				this.Auth.User.Salt = salt
				this.Auth.User.Password = Md5(newPwd + salt)
				this.Orm.Update(&this.Auth.User)
			} else {
				code = 1
				msg = "新密码和重复新密码不一致！"
			}
		} else {
			code = 1
			msg = "原密码错误!！"
		}
		this.Data["json"] = map[string]interface{}{"code": code, "msg": msg}
		this.ServeJSON()
		this.StopRun()
	}
}

// 站点配置
func (this *SysController) Config() {
	names := [2]string{"title", "url"}

	if this.Ctx.Input.IsPost() {
		for _, name := range names {
			this.Orm.InsertOrUpdate(&models.Config{Name: name, Value: this.GetString(name, "")})
		}
		this.Data["json"] = map[string]interface{}{"code": 0, "msg": ""}
		this.ServeJSON()
		this.StopRun()
	} else {
		configList := make(map[string]string)
		for _, name := range names {
			configList[name] = ""
		}
		var maps []orm.Params
		this.Orm.QueryTable(new(models.Config)).Values(&maps)
		for _, item := range maps {
			configList[item["Name"].(string)] = item["Value"].(string)
		}
		this.Data["Configs"] = configList
	}
}

// 后台首页
func (this *SysController) Home() {}

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

// 文章管理
func (this *SysController) Articles() {}

// 文章列表
func (this *SysController) ArticleList() {
	page, _ := this.GetInt("page", 1)
	limit, _ := this.GetInt("limit", 30)
	offset := (page - 1) * limit
	cnt, _ := this.Orm.QueryTable(new(models.Article)).Count()
	var list []models.Article
	this.Orm.QueryTable(new(models.Article)).RelatedSel().Limit(limit, offset).All(&list)
	//	var maps []orm.Params
	//	this.Orm.QueryTable(new(models.Article)).Limit(limit, offset).Values(&maps)
	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "count": cnt, "data": list}
	this.ServeJSON()
	this.StopRun()
}

// 文章详情(添加or编辑)
func (this *SysController) Article() {
	var categories []orm.Params
	this.Orm.QueryTable(new(models.Category)).Values(&categories)

	id, _ := this.GetInt(":id", 0)
	var article models.Article
	if id == 0 {
		article.Id = 0
		article.Category = new(models.Category)
	} else {
		article.Id = id
		this.Orm.Read(&article)
	}

	this.Data["Categories"] = categories
	this.Data["Article"] = article
}

// 保存文章
func (this *SysController) ArticleSave() {
	var article models.Article
	id, _ := this.GetInt("id", 0)
	categoryId, _ := this.GetInt("category_id", 0)
	if id == 0 {
		article.User = &this.Auth.User
	} else {
		article.Id = id
		this.Orm.Read(&article)
	}

	article.Title = this.GetString("title")
	article.Info = this.GetString("info")
	article.Content = this.GetString("content")
	article.ReleaseTime = StrToLocationTime(this.GetString("release_time"))

	var cat models.Category
	cat.Id = categoryId
	article.Category = &cat

	if id == 0 {
		this.Orm.Insert(&article)
	} else {
		this.Orm.Update(&article)
	}

	this.Data["json"] = map[string]interface{}{"code": 0, "msg": "", "data": article.Id}
	this.ServeJSON()
	this.StopRun()
}

// 删除文章
func (this *SysController) ArticleDel() {
	id, _ := this.GetInt("Id")
	this.Orm.Delete(&models.Article{Id: id})

	this.Data["json"] = map[string]interface{}{"code": 0, "msg": ""}
	this.ServeJSON()
	this.StopRun()
}

// 上传接口（ueditor）
// 返回数据格式：
// {
//     "state": "SUCCESS",
//     "url": "upload/demo.jpg",
//     "title": "demo.jpg",
//     "original": "demo.jpg"
// }
func (this *SysController) Upload() {
	file, fileHead, err := this.GetFile("upfile")
	defer file.Close()
	result := make(map[string]interface{})
	if err == nil {
		result["title"] = fileHead.Filename
		result["original"] = fileHead.Filename
		result["url"] = ""
		result["state"] = "FAIL"

		cPath, err := getCurrentPath()
		if err == nil {
			fName := "/static/img/upload/" + Md5(string(this.Auth.Id)+time.Now().String()+fileHead.Filename) + path.Ext(fileHead.Filename)
			err := this.SaveToFile("upfile", cPath+fName)
			if err == nil {
				result["state"] = "SUCCESS"
				result["url"] = fName
			} else {
				result["message"] = err.Error()
			}
		} else {
			result["message"] = err.Error()
		}
	} else {
		result["message"] = err.Error()
	}
	this.Data["json"] = result
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

// 时间转换
func StrToLocationTime(strTime string) time.Time {
	loc, _ := time.LoadLocation("Local")
	t, _ := time.ParseInLocation("2006-01-02 15:04:05", strTime, loc)
	return t
}

// 获取当前目录
func getCurrentPath() (string, error) {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		return "", err
	}
	path, err := filepath.Abs(file)
	if err != nil {
		return "", err
	}
	i := strings.LastIndex(path, "/")
	if i < 0 {
		i = strings.LastIndex(path, "\\")
	}
	if i < 0 {
		return "", errors.New(`error: Can't find "/" or "\".`)
	}
	return string(path[0 : i+1]), nil
}

// 生成随机字符串
func GetRandomString(l int) string {
	str := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	bytes := []byte(str)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
