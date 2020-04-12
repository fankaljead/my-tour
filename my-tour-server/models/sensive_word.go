package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// SensiveWord 敏感词
type SensiveWord struct {
	Id         int64     `json:"id"`
	KeyWord    string    `json:"key_word"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(SensiveWord))
	AddATable("sensive_word")
}
