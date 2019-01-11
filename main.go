package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "gobook/routers"
)

func init() {
	dbip := beego.AppConfig.String("dbip")
	dbport := beego.AppConfig.String("dbport")
	dbname := beego.AppConfig.String("dbname")
	dbuser := beego.AppConfig.String("dbuser")
	dbpassword := beego.AppConfig.String("dbpassword")
	dbMaxIdleConns, _ := beego.AppConfig.Int("dbMaxIdleConns")
	dbMaxOpenConns, _ := beego.AppConfig.Int("dbMaxOpenConns")

	orm.RegisterDriver("mysql", orm.DRMySQL)
	dataSourceName := dbuser + ":" + dbpassword + "@tcp(" + dbip + ":" + dbport + ")/" + dbname + "?charset=utf8&loc=Local"
	orm.RegisterDataBase("default", "mysql", dataSourceName)
	orm.SetMaxIdleConns("default", dbMaxIdleConns)
	orm.SetMaxOpenConns("default", dbMaxOpenConns)

	orm.RunSyncdb("default", false, true)
}

func main() {
	orm.Debug = true

	beego.BConfig.WebConfig.Session.SessionOn = true
	beego.BConfig.WebConfig.Session.SessionGCMaxLifetime = 60 * 30
	beego.BConfig.WebConfig.Session.SessionName = "gobooksessionID"

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.Run()
}
