package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type TransportWay struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Cover        string    `json:"cover"`
	CreateTime   time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	CreateUserId int64     `json:"create_user_id"`
}

func init() {
	orm.RegisterModel(new(TransportWay))
	AddATable("transport_way")
}
