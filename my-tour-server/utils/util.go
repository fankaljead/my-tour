package utils

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", "zxh:212kawhi@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
}
