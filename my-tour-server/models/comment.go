package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

// ScenarySpotCommet
type ScenarySpotCommet struct {
	Id            int64   `json:"id"`
	Title         string  `json:"title"`
	Content       string  `json:"content"`
	UserId        int64   `json:"user_id"`
	ScenarySpotId int64   `json:"scenary_spot_id"`
	StarNumber    int     `json:"star_number"`
	Score         float64 `json:"score"`
}

// travel_note 旅游记录评论
type TravelNoteComment struct {
	Id           int64  `json:"id"`
	Content      string `json:"content"`
	TravelNoteId int64  `json:"travel_note_id"`
	UserId       int64  `json:"user_id"`
	CreateTime   string `json:"create_time"`
}

func init() {
	orm.RegisterModel(new(ScenarySpotCommet), new(TravelNoteComment))
	AddATable("scenary_spot_comment", "travel_note_comment")
}

// AddATravelNoteCommet 新增一条评论
func AddATravelNoteCommet(content string, travel_note_id, user_id int64) (int64, error) {
	travel_note_comment := TravelNoteComment{
		UserId:       user_id,
		Content:      content,
		TravelNoteId: travel_note_id,
		CreateTime:   time.Now().String()}

	o := orm.NewOrm()
	id, err := o.Insert(&travel_note_comment)

	return id, err
}

// DeleveTravelNoteComment 删除一条评论
func DeleveTravelNoteComment(travel_note_comment_id int64) (num int64) {
	o := orm.NewOrm()
	num, _ = o.Delete(&TravelNoteComment{Id: travel_note_comment_id})
	return
}

// GetTravelNoteComments 获取评论
func GetTravelNoteComments(page_size, page_index, travel_note_id int64) map[string]interface{} {
	travel_note_comments := make(map[string]interface{})

	o := orm.NewOrm()

	// comments := make([]TravelNoteComment, 10)
	var comments []TravelNoteComment

	start := page_size * (page_index - 1)
	end := page_size * page_index

	number, err := o.Raw("select * from travel_note_comment " + " where travel_note_id=" + strconv.FormatInt(travel_note_id, 10) + " order by id desc" + " limit " +
		strconv.FormatInt(start, 10) +
		"," + strconv.FormatInt(end, 10)).QueryRows(&comments)

	if err == nil {
		travel_note_comments["comments"] = comments
	}

	travel_note_comments["number"] = number
	travel_note_comments["page_size"] = page_size
	travel_note_comments["page_index"] = page_index

	return travel_note_comments
}

func UpdateTravelNoteComment(comment_id int64, content string) (num int64, err error) {
	o := orm.NewOrm()

	fmt.Printf("comment_id: %d\n", comment_id)
	num, err = o.QueryTable("travel_note_comment").Filter("id", comment_id).Update(orm.Params{
		"content": content,
	})
	return
}
