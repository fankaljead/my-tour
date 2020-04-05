package models

import (
	"fmt"
	"time"

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

func GetUserInfo(user_id string) (user_info UserInfo, err error) {
	o := orm.NewOrm()
	sql := "select * from " + user_info.TableName() + " where user_id=?"
	err = o.Raw(sql, user_id).QueryRow(&user_info)

	return
}

func GetAllUserInfo() (user_infos []UserInfo, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("select * from user_info").QueryRows(&user_infos)
	return
}

func SetUserHeadIcon(user_id string, head_icon string) (num int64, err error) {
	o := orm.NewOrm()
	update_time := time.Now().String()
	res, err := o.Raw("update user_info set head_icon=? , update_time = ? where user_id=?", head_icon, update_time, user_id).Exec()
	if err == nil {
		num, err = res.RowsAffected()
	}
	return
}

func SetUserBackgroundImage(user_id string, background_image string) (num int64, err error) {
	o := orm.NewOrm()
	update_time := time.Now().String()
	res, err := o.Raw("update user_info set background_image=? , update_time = ? where user_id=?", background_image, update_time, user_id).Exec()
	if err == nil {
		num, err = res.RowsAffected()
	}
	return
}
