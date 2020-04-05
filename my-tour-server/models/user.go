package models

import (
	"strconv"

	"github.com/astaxie/beego/orm"
	// _ "github.com/go-sql-driver/mysql"
)

const (
	// 登录成功
	LOGIN_SUCCESS = iota
	// 登录失败 用户不存在
	LOGIN_USER_NOT_EXIST
	// 登录失败 密码错误
	LOGIN_PASSWORD_NOT_MATCH
)

func init() {

	// orm.RegisterDriver("mysql", orm.DRMySQL)
	// orm.RegisterDataBase("default", "mysql", "zxh:212kawhi@tcp(127.0.0.1:3306)/test?charset=utf8", 30)
	orm.RegisterModel(new(User))
	// create table
	// orm.RunSyncdb("default", true, true)
}

type User struct {
	Id       int64  `form:"-"` //`orm:"pk"`
	Username string `form:"username"`
	Password string `form:"password"`
	// Id       int64
	// Username string
	// Password string
}

func AddUser(u User) int64 {
	o := orm.NewOrm()
	id, err := o.Insert(&u)
	if err != nil {

	}
	return id
}

func GetUser(uid string) (u *User, err error) {
	o := orm.NewOrm()
	err = o.Raw("SELECT id, username,password FROM user WHERE id = ?", uid).QueryRow(&u)
	return
}

func GetAllUsers() (users []User, num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Raw("SELECT id, username,password FROM user ").QueryRows(&users)
	return
}

func UpdateUser(uid string, uu *User) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.QueryTable("user").Filter("id", uu.Id).Update(orm.Params{
		"username": uu.Username,
		"password": uu.Password,
	})
	return
}

func Login(username, password string) int64 {
	o := orm.NewOrm()

	var user User
	err := o.Raw("SELECT id, password FROM user WHERE username = ?", username).QueryRow(&user)

	if err != nil {
		return LOGIN_USER_NOT_EXIST
	}
	if password == user.Password {
		return user.Id
	}
	return LOGIN_PASSWORD_NOT_MATCH
}

func DeleteUser(uid string) (num int64, err error) {
	o := orm.NewOrm()
	id, _ := strconv.ParseInt(uid, 10, 64)
	num, err = o.Delete(&User{Id: id})
	return
}
