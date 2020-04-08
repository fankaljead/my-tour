package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Fan struct {
	Id         int64 `json:"id"`
	FollowerId int64 `json:"follower_id"`
	UserId     int64 `json:"user_id"`
	Time       string
}

func init() {
	orm.RegisterModel(new(Fan))
}

func FollowUser(user_id int64, follow_id int64) (int64, error) {
	fan := Fan{UserId: user_id, FollowerId: follow_id}
	fan.Time = time.Now().String()
	o := orm.NewOrm()
	id, err := o.Insert(&fan)

	return id, err
}

func UnfollowUser(user_id int64, follow_id int64) int64 {
	fan := Fan{UserId: user_id, FollowerId: follow_id}
	o := orm.NewOrm()
	if num, err := o.Delete(&fan); err == nil {
		return num
	}
	return -1
}

func GetAllFollowers(user_id string) (followers_info map[string]interface{}) {
	o := orm.NewOrm()

	var followers []Fan
	num, err := o.Raw("SELECT id, follower_id FROM user WHERE id = ?", user_id).QueryRows(&followers)
	if err == nil {
		followers_info["followers"] = followers
	}
	followers_info["numeber"] = num
	return
}
