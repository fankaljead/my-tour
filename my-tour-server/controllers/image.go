package controllers

import (
	"fmt"
	"strconv"
	"strings"

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
	user_id := GetCurrentSessionUserIdString()

	if err != nil {
		c.Data["json"] = 0
		// log.Fatal("getfile err ", err)
		fmt.Println(err)
	} else {
		image_name := user_id + h.Filename
		c.SaveToFile("image_name", beego.AppConfig.String("uploadStaticDir")+"/"+image_name) // 保存位置在 static/upload, 没有文件夹要先创建

		id, _ := models.UploadImage(image_name)

		// m := make(map[string]int64)

		// m["image_id"] = id
		c.Data["json"] = id
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

// @Title Get
// @Description get image by uid
// @Param	uid		path 	int	true		"The key for image"
// @Success 200 {object} models.Image
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *ImageController) Get() {
	uid, _ := u.GetInt64(":uid")
	image := models.GetImageById(uid)
	u.Data["json"] = image
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users information
// @Param	image_ids		path 	string	true		"The ids for images"
// @Success 200 {object} []models.Image
// @router /getImagesByIds [get]
func (u *ImageController) GetImagesByIds() {

	image_ids := u.GetString("image_ids")

	idss := strings.Split(image_ids, ":")

	ids := []int64{}

	for _, i := range idss {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		ids = append(ids, int64(j))
	}

	images := models.GetImagesByIds(ids)

	u.Data["json"] = images
	u.ServeJSON()
}
