package models

import (
	"time"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	// 注册mysql orm
	dbuser := beego.AppConfig.String("mysqluser")
	dbpwd := beego.AppConfig.String("mysqlpwd")
	dbhost := beego.AppConfig.String("mysqlhost")
	dbport := beego.AppConfig.String("mysqlport")
	dbname := beego.AppConfig.String("mysqldbname")
	conStr := dbuser + ":" + dbpwd + "@tcp(" + dbhost + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"

	orm.DefaultTimeLoc = time.Local
	orm.RegisterDataBase("default", "mysql", conStr)

	prefix := beego.AppConfig.String("mysqlprefix")
	orm.RegisterModelWithPrefix(prefix, new(User), new(Category), new(Article), new(Config))
}

// 用户
type User struct {
	Id            int
	Username      string
	Password      string
	Salt          string
	State         int
	LastTime      time.Time
	LastIp        string
	LoginErrTimes int
	Article       []*Article `orm:"reverse(many)"`
}

// 分类
type Category struct {
	Id      int
	Name    string
	Sort    int
	Created time.Time  `orm:"auto_now_add;type(datetime)"`
	Updated time.Time  `orm:"auto_now;type(datetime)"`
	Article []*Article `orm:"reverse(many)"`
}

// 文章
type Article struct {
	Id int
	//UserId int
	//CategoryId  int
	Title       string
	Content     string
	Info        string
	ReleaseTime time.Time
	Created     time.Time `orm:"auto_now_add;type(datetime)"`
	Updated     time.Time `orm:"auto_now;type(datetime)"`
	User        *User     `orm:"rel(fk)"`
	Category    *Category `orm:"rel(fk)"`
}

// 配置
type Config struct {
	Name  string `orm:"pk"`
	Value string
}
