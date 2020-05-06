package controllers

import (
	"log"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

// Operations about UserInfo
type ImageController struct {
	beego.Controller
}

// @Title Upload A Image
// @Description upload image
// @Param	image_name		path 	file	true		"The key for staticblock"
// @Success 200 {int} models.Image.Id
// @Failure 403 file is empty
// @Content-Type multipart/form-data
// @router / [post]
func (c *ImageController) Post() {
	f, h, err := c.GetFile("image_name")
	// user_id := c.GetString(":user_id")
	user_id := "1"

	if err != nil {
		c.Data["json"] = 0
		log.Fatal("getfile err ", err)
	} else {
		image_name := beego.AppConfig.String("uploadStaticDir") + "/" + user_id + h.Filename
		c.SaveToFile("uploadname", image_name) // 保存位置在 static/upload, 没有文件夹要先创建

		num, _ := models.UploadImage(image_name)

		c.Data["json"] = num
	}
	defer f.Close()
	c.ServeJSON()
}

// @Title Delete Image
// @Description delete image
// @Param	uid		path 	string	true		"The image uid you want to delete"
// @Success 200 {int} 1 delete success!
// @Failure 403 0 uid is empty
// @router /:uid [delete]
func (u *ImageController) Delete() {
	image_id, err := u.GetInt64(":uid")
	num, err := models.DeleteImage(image_id)
	if err != nil {
		u.Data["json"] = 0
	}
	u.Data["json"] = num
	u.ServeJSON()
}
