package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// Tag 标签
type Tag struct {
	Id           int64     `json:"id"`
	Name         string    `json:"name"`
	Description  string    `json:"description"`
	CreateUserId int64     `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

// TagsTo 标签对应
type TagsTo struct {
	Id         int64     `json:"id"`
	ToId       int64     `json:"to_id"`
	TagId      int64     `json:"tag_id"`
	ToTypeId   int64     `json:"to_type_id"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(Tag), new(TagsTo))
	AddATable("tag", "tags_to")
}
