package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Image struct {
	Id         int64     `json:"id"`
	Source     string    `json:"source"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func init() {

	orm.RegisterModel(new(Image))
	AddATable("image")
}

// UploadImage 上传图片
func UploadImage(image_name string) (image_id int64, err error) {

	image := Image{
		Source: image_name,
	}
	o := orm.NewOrm()
	image_id, err = o.Insert(&image)

	return
}

// DeleteImage 删除图片
func DeleteImage(image_id int64) (num int64, err error) {
	o := orm.NewOrm()
	num, err = o.Delete(&Image{Id: image_id})
	return
}

// GetImagesByIds 获取多张图片
func GetImagesByIds(ids []int64) (images []Image) {

	o := orm.NewOrm()

	for _, id := range ids {
		image := Image{Id: id}
		o.Read(&image)
		images = append(images, image)
	}

	return
}

// GetImageById 获取单张图片
func GetImageById(id int64) (image Image) {

	o := orm.NewOrm()
	image = Image{Id: id}
	o.Read(&image)

	return
}
