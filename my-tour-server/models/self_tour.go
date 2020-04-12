package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type SelfTour struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	Content      string    `json:"content"`
	Cover        string    `json:"cover"`
	CreateUserId int64     `json:"create_user_id"`
	CreateUptime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

type SelfTourDays struct {
	Id            int64     `json:"id"`
	Title         string    `json:"title"`
	Descirption   string    `json:"descirption"`
	DayNumber     int       `json:"day_number"`
	SelfTourId    int64     `json:"self_tour_id"`
	ScenarySpotId int64     `json:"scenary_spot_id"`
	CreateUserId  int64     `json:"create_user_id"`
	CreateUptime  time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(SelfTour), new(SelfTourDays))
	AddATable("self_tour", "self_tour_days")
}
