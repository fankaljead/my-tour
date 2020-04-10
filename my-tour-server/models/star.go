package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Star struct {
	Id          int64     `json:"id"`
	StarTypeId  int64     `json:"star_type_id"`
	StartUserId int64     `json:"start_user_id"`
	CreateTime  time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

type StarType struct {
	Id          int64 `json:"id"`
	TypeTableId int64 `json:"type_table_id"`
}

func init() {
	orm.RegisterModel(new(Star), new(StarType))
	AddATable("star")
	AddATable("star_type")
}

func AddAStar(start_type_id, star_user_id int64) (int64, error) {
	star := Star{
		StarTypeId:  start_type_id,
		StartUserId: star_user_id,
	}

	o := orm.NewOrm()

	return o.Insert(&star)
}

func CancelAStar(star_id int64) (int64, error) {
	star := Star{
		Id: star_id,
	}

	o := orm.NewOrm()

	return o.Delete(&star)
}
