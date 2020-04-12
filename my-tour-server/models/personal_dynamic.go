package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// PersonalDynamic 个人动态
type PersonalDynamic struct {
	Id           int64     `json:"id"`
	Type         string    `json:"type"`
	Content      string    `json:"content"`
	Draft        bool      `json:"draft"`
	CreateUserId int64     `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	PublishTime  time.Time `json:"publish_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(PersonalDynamic))
	AddATable("personal_dynamic")
}
