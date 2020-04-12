package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// LoginLog 登录日志
type LoginLog struct {
	Id        int64     `json:"id"`
	UserIp    string    `json:"user_ip"`
	UserId    int64     `json:"user_id"`
	LoginTime time.Time `json:"login_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(LoginLog))
	AddATable("login_log")
}
