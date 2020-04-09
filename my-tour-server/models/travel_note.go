package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type TravelNote struct {
	Id                int64  `json:"travel_note_id"`
	AuthorId          int64  `json:"author_id"`
	Title             string `json:"title"`
	Place             string `json:"place"`
	Content           string `json:"content"`
	Cover             string `json:"cover"`
	Draft             bool   `json:"draft"`
	TravelNoteTopicId int64  `json:"travel_note_topic_id"`
	PublishTime       string `json:"publish_time"`
	CreateTime        string `json:"creat_time"`
}

type TravelNoteTopic struct {
	Id           int64
	CreateUserId int64
	Name         string
	Description  string
	CreateTime   string
}

func init() {
	orm.RegisterModel(new(TravelNote))
}

func AddTravelNote(title string, content string, cover string,
	place string, user_id int64, publishTime string) (int64, error) {

	travel_note := TravelNote{
		AuthorId:    user_id,
		Title:       title,
		Content:     content,
		Cover:       cover,
		PublishTime: publishTime,
		CreateTime:  time.Now().String(),
		Place:       place}

	o := orm.NewOrm()
	id, err := o.Insert(&travel_note)

	return id, err
}

func UpdateTravelNote(travelNote *TravelNote) int64 {
	o := orm.NewOrm()
	status, err := o.Update(travelNote)
	if err == nil {
		return status
	}
	return -1
}

func DeleveTravelNote(travelNoteId int64) (num int64) {
	o := orm.NewOrm()
	num, _ = o.Delete(&TravelNote{Id: travelNoteId})
	return
}

func GetTravelNote(travelNoteId int64) (TravelNote, error) {
	o := orm.NewOrm()
	travelNote := TravelNote{
		Id: travelNoteId}
	err := o.Read(&travelNote)
	return travelNote, err
}

func GetAllUserTravelNoteInfo(user_id string) map[string]interface{} {
	o := orm.NewOrm()

	user_travel_note_infos := make(map[string]interface{})

	user_travel_note_infos["user_id"] = user_id

	var res []struct {
		Id         int64  `json:"id"`
		Title      string `json:"title"`
		Cover      string `json:"cover"`
		CreateTime string `json:"create_time"`
	}
	nums, err := o.Raw("select id,title,cover,create_time from travel_note where author_id=" + user_id).QueryRows(&res)

	if err == nil {
		user_travel_note_infos["travelNoteInfos"] = res
	}

	user_travel_note_infos["number"] = nums
	return user_travel_note_infos
}

func SetTravelNoteDraft(travel_note_id int64) int64 {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE travel_note SET draft = false where id=?", travel_note_id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num
	}

	return -1
}

func PublishTravelNoteDraft(travel_note_id int64) int64 {
	o := orm.NewOrm()
	res, err := o.Raw("UPDATE travel_note SET draft = true where id=?", travel_note_id).Exec()
	if err == nil {
		num, _ := res.RowsAffected()
		return num
	}

	return -1
}
