package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Star struct {
	Id          int64     `json:"id"`
	StarType    int64     `json:"star_type"`
	StartUserId int64     `json:"start_user_id"`
	CreateTime  time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

type StarType struct {
	Id            int64  `json:"id"`
	TypeTableName string `json:"type_table_name"`
}

func init() {
	orm.RegisterModel(new(Star), new(StarType))
	AddATable("star")
	AddATable("star_type")
}
