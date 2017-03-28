package main

import (
	_ "env/models"
	_ "env/routers"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
)

var (
	DB_TYPE = beego.AppConfig.String("DB_TYPE")
	DB_NAME = beego.AppConfig.String("DB_NAME")
	DB_USER = beego.AppConfig.String("DB_USER")
	DB_PASSWD = beego.AppConfig.String("DB_PASSWD")
)

func init() {
	var register = DB_USER + ":" + DB_PASSWD + "@/" + DB_NAME + "?charset=utf8&loc=Local"
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", DB_TYPE, register, 30)
	orm.RunSyncdb("default", true, false)
}

func main() {
	beego.Run()
}
