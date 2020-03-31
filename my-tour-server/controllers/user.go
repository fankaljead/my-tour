package controllers

import (
	"fmt"
	"strconv"

	"gitlab.com/fankaljead/my-tour/my-tour-server/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	var user models.User
	// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	user.Username = u.GetString("username")
	user.Password = u.GetString("password")
	fmt.Println("user information: ")
	fmt.Println(user)
	uid := models.AddUser(user)
	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router / [get]
func (u *UserController) GetAll() {
	users, num, _ := models.GetAllUsers()
	users_info := make(map[string]interface{})

	users_info["users"] = users
	users_info["length"] = num
	u.Data["json"] = users_info
	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	uid := u.GetString(":uid")
	if uid != "" {
		user, err := models.GetUser(uid)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = user
		}
	}
	u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	fmt.Println("in update method")
	uid := u.GetString("id")
	if uid != "" {
		var user models.User
		// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
		username := u.GetString("username")
		password := u.GetString("password")

		fmt.Println("uid")
		fmt.Println(uid)

		if uid != "" {
			user.Id, _ = strconv.ParseInt(uid, 10, 64)
		}

		if username != "" {
			user.Username = username
		}
		if password != "" {
			user.Password = password
		}

		fmt.Println(user)

		uu, err := models.UpdateUser(uid, &user)
		if err != nil {
			u.Data["json"] = err.Error()
		} else {
			u.Data["json"] = uu
		}
	}
	u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {int} 1 delete success!
// @Failure 403 0 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	uid := u.GetString(":uid")
	num, err := models.DeleteUser(uid)
	if err != nil {
		u.Data["json"] = 0
	}
	u.Data["json"] = num
	u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (u *UserController) Login() {

	username := u.GetString("username")
	password := u.GetString("password")
	fmt.Println("Login in ")
	fmt.Println(username)
	switch models.Login(username, password) {
	case models.LOGIN_SUCCESS:
		u.Data["json"] = "login success"
		break
	case models.LOGIN_PASSWORD_NOT_MATCH:
		u.Data["json"] = "password does not match"
		break
	case models.LOGIN_USER_NOT_EXIST:
		u.Data["json"] = "user not exist"
		break
	}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
