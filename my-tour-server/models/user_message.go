package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type UserMessage struct {
	Id         int64     `json:"id"`
	Type       int64     `json:"type"`
	Content    string    `json:"content"`
	SenderId   int64     `json:"sender_id"`
	ReceiverId int64     `json:"receiver_id"`
	IsReceive  bool      `json:"is_receive"`
	SendTime   time.Time `json:"send_time"`
}
type UserMessageType struct {
	Id         int64     `json:"id"`
	Name       string    `json:"nme"`
	CreateTime time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

func init() {
	orm.RegisterModel(new(UserMessage), new(UserMessageType))
	AddATable("user_message")
}
