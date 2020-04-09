package controllers

import (
	"fmt"
	"strconv"

	"gitlab.com/fankaljead/my-tour/my-tour-server/models"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/session"
)

// Global Sessions
var GlobalSessions *session.Manager

func init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  3600,
		ProviderConfig:  "./tmp",
	}
	GlobalSessions, _ = session.NewManager("memory", sessionConfig)
	go GlobalSessions.GC()
}

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

	r := u.Ctx.Request
	w := u.Ctx.ResponseWriter
	sess, _ := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)

	fmt.Println("Login in ")
	fmt.Println(username)
	user_id := models.Login(username, password)
	fmt.Println("user_id")
	fmt.Println(user_id)
	u.Data["json"] = user_id

	sess.Set("user_id", user_id)

	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"

	r := u.Ctx.Request
	w := u.Ctx.ResponseWriter
	GlobalSessions.SessionDestroy(w, r)

	u.ServeJSON()
}

func GetCurrentSessionUserIdInt64() int64 {
	return 1
}

func GetCurrentSessionUserIdString() string {
	return "1"
}
