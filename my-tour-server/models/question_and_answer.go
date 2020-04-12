package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

// QuestionAndAnswer 问答
type QuestionAndAnswer struct {
	Id           int64     `json:"id"`
	Title        string    `json:"title"`
	Content      string    `json:"content"`
	Draft        bool      `json:"draft"`
	CreateUserId int64     `json:"create_user_id"`
	CreateTime   time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
}

// QAndAAnswer 问答回答
type QAndAAnswer struct {
	Id           int64     `json:"id"`
	QuestionId   int64     `json:"question_id"`
	Content      string    `json:"content"`
	AnswerUserId int64     `json:"answer_user_id"`
	CreateTime   time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	StarNumber   int       `json:"star_number"`
}

// QAndAAnswerComment 问答回答评论
type QAndAAnswerComment struct {
	Id                         int64     `json:"id"`
	AnswerId                   int64     `json:"answer_id"`
	QAndAAnswerAnswerCommentId int64     `json:"q_and_a_answer_comment_id"`
	Content                    string    `json:"content"`
	AnswerUserId               int64     `json:"answer_user_id"`
	CreateTime                 time.Time `json:"create_time" orm:"auto_now_add;type(datetime)"`
	StarNumber                 int       `json:"star_number"`
}

func init() {
	orm.RegisterModel(new(QuestionAndAnswer), new(QAndAAnswer), new(QAndAAnswerComment))
	AddATable("question_and_answer", "q_and_a_answer", "q_and_a_answer_comment")
}
