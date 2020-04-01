package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

// Operations about UserInfo
type User_InfoController struct {
	beego.Controller
}

// @Title UpdateUserInfo
// @Description create user information
// @Param	body		body 	models.UserInfo	true		"body for user_info content"
// @Success 200 {int} models.UserInfo.Id
// @Failure 403 body is empty
// @router / [put]
func (u *User_InfoController) Put() {
	// var user_info *models.UserInfo
	user_info := new(models.UserInfo)
	// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// user_info. = u.GetString("username")
	// user_id, _ := strconv.ParseInt(u.GetString("user_id"), 10, 64)
	user_id, _ := strconv.ParseInt(u.GetString("user_id"), 10, 64)
	tel := u.GetString("tel")
	birthday := u.GetString("birthday")
	qq := u.GetString("qq")
	email := u.GetString("email")
	location := u.GetString("location")
	personalized_signature := u.GetString("personalized_signature")
	wechat := u.GetString("wechat")
	background_image := u.GetString("background_image")
	head_icon := u.GetString("head_icon")

	// user_id, _ := u.GetInt64("user_id")

	user_info.UserId = user_id
	if tel != "" {
		user_info.Tel = tel
	}
	if qq != "" {
		user_info.Qq = qq
	}
	if birthday != "" {
		user_info.Birthday = birthday
	}
	if email != "" {
		user_info.Email = email
	}
	if location != "" {
		user_info.Location = location
	}
	if wechat != "" {
		user_info.Wechat = wechat
	}
	if personalized_signature != "" {
		user_info.PersonalizedSignature = personalized_signature
	}
	if background_image != "" {
		user_info.BackgroundImage = background_image
	}
	if head_icon != "" {
		user_info.HeadIcon = head_icon
	}

	// user_info.Tel = u.GetString("tel")

	uid := models.UpdateUserInfo(user_info)

	u.Data["json"] = uid
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users information
// @Success 200 {object} models.UserInfo
// @router / [get]
func (u *User_InfoController) GetAll() {
	u.Data["json"] = ""
	u.ServeJSON()
}

// @Title Get
// @Description get user_info by user_id
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserInfo
// @Failure 403 :uid is empty
// @router /:user_id [get]
func (u *User_InfoController) Get() {
	user_id := u.GetString(":user_id")
	if user_id != "" {
		user_info, _ := models.GetUserInfo(user_id)
		u.Data["json"] = user_info
	}
	u.ServeJSON()
}
