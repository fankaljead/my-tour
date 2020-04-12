package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type Photo struct {
	Id                    int64     `json:"id"`
	Type                  string    `json:"type"`
	Src                   string    `json:"src"`
	GuidelineOrLocationId int64     `json:"guideline_or_location_id"`
	UploadUserId          int64     `json:"upload_user_id"`
	UploadTime            time.Time `json:"upload_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Photo))
	AddATable("photo")
}
