package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Guideline struct {
	Id              int64     `json:"id"`
	FromWhere       string    `json:"from_where"`
	ToWhere         string    `json:"to_where"`
	TransportWayId  int64     `json:"transport_way_id"`
	PersonalThought string    `json:"personal_thought"`
	CreateUserId    int64     `json:"create_user_id"`
	CreateTime      time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Guideline))
	AddATable("guideline")
}
