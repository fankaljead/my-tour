package main

import (
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
	_ "gitlab.com/fankaljead/my-tour/my-tour-server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	sqlDriver := beego.AppConfig.String("sqlDriver")
	sqlconn := beego.AppConfig.String("sqlconn")
	name := "default"
	force := false
	verbose := true

	orm.RegisterDriver(sqlDriver, orm.DRMySQL)
	orm.RegisterDataBase(name, sqlDriver, sqlconn, 30)
	// create table
	orm.RunSyncdb(name, force, verbose)
}
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
		models.AutoCreateTabel()
	}
	beego.Run()
}
