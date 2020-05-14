package main

import (
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
	_ "gitlab.com/fankaljead/my-tour/my-tour-server/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
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
	// beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
	// 	AllowOrigins:     []string{"https://*.foo.com"},
	// 	AllowMethods:     []string{"PUT", "PATCH"},
	// 	AllowHeaders:     []string{"Origin"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true}))
	beego.BConfig.WebConfig.DirectoryIndex = true
	beego.SetStaticPath(beego.AppConfig.String("statiImageUrl"),
		beego.AppConfig.String("uploadStaticDir"))

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		//AllowAllOrigins:  true,
		AllowMethods:     []string{"*"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "content-type"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Methods", "Access-Control-Allow-Origin", "content-type"},
		AllowCredentials: true,
		AllowOrigins:     []string{"http://10.*.*.*:*", "http://localhost:*", "http://127.0.0.1:*", "*"},
	}))
	beego.Run()
}
