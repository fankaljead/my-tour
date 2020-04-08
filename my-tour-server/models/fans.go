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
	o := orm.NewOrm()

	res, err := o.Raw("delete from fan where user_id =? and follower_id=?", user_id, follow_id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num
	}

	return -1
}

func GetAllFollowers(user_id string) (followers_info map[string]interface{}) {
	o := orm.NewOrm()

	followers_info = make(map[string]interface{})

	var followers []Fan
	num, err := o.Raw("SELECT id, follower_id,user_id,time FROM fan WHERE user_id = ?", user_id).QueryRows(&followers)
	if err == nil {
		followers_info["followers"] = followers
	}
	followers_info["numeber"] = num

	return
}
