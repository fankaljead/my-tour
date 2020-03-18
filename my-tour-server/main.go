package main

import (
	"fmt"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
	_ "gitlab.com/fankaljead/my-tour/my-tour-server/routers"
)

func init() {
	err := orm.RegisterDataBase("default", "mysql", "zxh:212kawhi@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	fmt.Println(err)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
