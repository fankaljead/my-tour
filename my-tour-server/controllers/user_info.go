package controllers

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

// Operations about UserInfo
type UserInfoController struct {
	beego.Controller
}

// @Title UpdateUserInfo
// @Description create user information
// @Param	body		body 	models.UserInfo	true		"body for user_info content"
// @Success 200 {int} models.UserInfo.Id
// @Failure 403 body is empty
// @router / [put]
func (u *UserInfoController) Put() {
	// var user_info *models.UserInfo
	user_info := new(models.UserInfo)
	// json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	// user_info. = u.GetString("username")
	// user_id, _ := strconv.ParseInt(u.GetString("user_id"), 10, 64)
	// user_id := utils.GetSessionUserId(u)
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

	user_info.UpdateTime = time.Now().String()

	// user_info.Tel = u.GetString("tel")

	uid := models.UpdateUserInfo(user_info)

	u.Data["json"] = uid
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users information
// @Success 200 {object} models.UserInfo
// @router / [get]
func (u *UserInfoController) GetAll() {
	data := make(map[string]interface{})

	user_infos, num, _ := models.GetAllUserInfo()

	data["user_infos"] = user_infos
	data["length"] = num

	u.Data["json"] = data
	u.ServeJSON()
}

// @Title Get
// @Description get user_info by user_id
// @Param	user_id		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.UserInfo
// @Failure 403 :uid is empty
// @router /:user_id [get]
func (u *UserInfoController) Get() {
	user_id := u.GetString(":user_id")
	if user_id != "" {
		user_info, _ := models.GetUserInfo(user_id)
		u.Data["json"] = user_info
	}
	u.ServeJSON()
}

// @Title Set Head Icon
// @Description set icon for this user
// @Param	uploadname		path 	file	true		"The key for staticblock"
// @Success 200 {int} 1
// @Failure 403 file is empty
// @Content-Type multipart/form-data
// @router /set_icon [post]
func (c *UserInfoController) SetIcon() {
	f, h, err := c.GetFile("uploadname")
	// user_id := c.GetString(":user_id")
	user_id := "1"

	if err != nil {
		c.Data["json"] = 0
		log.Fatal("getfile err ", err)
	} else {
		head_icon := beego.AppConfig.String("uploadStaticDir") + "/" + user_id + h.Filename
		c.SaveToFile("uploadname", head_icon) // 保存位置在 static/upload, 没有文件夹要先创建

		num, _ := models.SetUserHeadIcon(user_id, head_icon)

		c.Data["json"] = num
	}
	defer f.Close()
	c.ServeJSON()
}

// @Title Set Background Image
// @Description set background image for this user
// @Param	uploadname		path 	file	true		"The key for staticblock"
// @Success 200 {int} 1
// @Failure 403 file is empty
// @Content-Type multipart/form-data
// @router /set_background_image [post]
func (c *UserInfoController) SetBackgroundImage() {
	f, h, err := c.GetFile("uploadname")
	// user_id := c.GetString(":user_id")
	// user_id := "1"

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	sess, _ := GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	user_id := fmt.Sprintf("%v", sess.Get("user_id"))

	fmt.Println("session user id")
	fmt.Println(user_id)

	if err != nil {
		c.Data["json"] = 0
		log.Fatal("getfile err ", err)
	} else {
		head_icon := beego.AppConfig.String("uploadStaticDir") + "/" + user_id + h.Filename
		c.SaveToFile("uploadname", head_icon) // 保存位置在 static/upload, 没有文件夹要先创建

		num, _ := models.SetUserBackgroundImage(user_id, head_icon)

		c.Data["json"] = num
	}
	defer f.Close()
	c.ServeJSON()
}
