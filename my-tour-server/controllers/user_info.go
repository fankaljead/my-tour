package controllers

import (
	"strconv"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

// Operations about User_Info
type User_InfoController struct {
	beego.Controller
}

// @Title CreateUser_Info
// @Description create user information
// @Param	body		body 	models.User_Info	true		"body for user_info content"
// @Success 200 {int} models.User_Info.Id
// @Failure 403 body is empty
// @router / [post]
func (u *User_InfoController) Post() {
	var user_info *models.User_Info
	// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// user_info. = u.GetString("username")
	user_id, _ := strconv.ParseInt(u.GetString("user_id"), 10, 64)
	user_info.User_id = user_id

	user_info.Tel = u.GetString("tel")

	uid := models.AddUserInfo(user_info)

	u.Data["json"] = map[string]int64{"uid": uid}
	u.ServeJSON()
}
