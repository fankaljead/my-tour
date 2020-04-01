package models

import (
	"fmt"

	"github.com/astaxie/beego/orm"
)

const (
	UNKNOWN = 0
	MALE    = 1
	FAMALE  = 2
)

func init() {

	orm.RegisterModel(new(UserInfo))
	// create table
	// orm.RunSyncdb("default", true, true)
}

type UserInfo struct {
	Id                    int64 `form:"-"` //`orm:"pk"`
	UserId                int64
	Tel                   string
	Birthday              string
	HeadIcon              string
	Qq                    string
	Email                 string
	Location              string
	PersonalizedSignature string
	BackgroundImage       string
	Wechat                string
	UpdateTime            string
}

func UpdateUserInfo(user_info *UserInfo) int64 {
	o := orm.NewOrm()
	// sql := "SELECT * FROM " + user_info.TableName() + " WHERE user_id = ?"
	// fmt.Println(sql)
	// var u UserInfo
	// err := o.Raw(sql, user_info.UserId).QueryRow(&u)

	// // if err != nil {
	// // 	return -2
	// // }
	// user_info.Id = u.Id

	// fmt.Printf("\n\nuser info id: %d\n", u.Id)
	// fmt.Printf("\n\nuser info user id: %d\n", u.UserId)
	// fmt.Println(u)
	// fmt.Println(user_info)
	// id, err := o.Update(user_info)
	id, err := o.InsertOrUpdate(user_info)

	if err != nil {
		fmt.Println(err)
		return -1
	}

	return id
}
func (u *UserInfo) TableName() string {
	return "user_info"
}

func GetUserInfo(user_id string) (user_info *UserInfo, err error) {
	o := orm.NewOrm()
	sql := "select * from " + user_info.TableName() + " where user_id="
	fmt.Println(sql)
	err = o.Raw(sql, user_id).QueryRow(&user_info)

	return
}
