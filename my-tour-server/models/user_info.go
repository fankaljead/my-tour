package models

import "github.com/astaxie/beego/orm"

const (
	UNKNOWN = 0
	MALE    = 1
	FAMALE  = 2
)

type User_Info struct {
	Id                     int64 `form:"-"` //`orm:"pk"`
	User_id                int64
	Tel                    string
	Birthday               string
	Head_Icon              string
	Qq                     string
	Email                  string
	Location               string
	Hobby                  int
	Personalized_Signature string
	Background_Image       string
	Wechat                 string
	Register_Time          string
}

func AddUserInfo(user_info *User_Info) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(user_info)
	if err != nil {
	}
	return id
}
