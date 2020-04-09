package models

import (
	"time"

	"github.com/astaxie/beego/orm"
)

type ScenarySpot struct {
	Id                    int64
	Name                  string
	ShortName             string
	Background            string
	Location              string
	Description           string
	EditTime              string
	Score                 float64
	ScenarySpotCatagoryId int64
}

func init() {
	orm.RegisterModel(new(ScenarySpot))
}

func AddScenarySpot(name, short_name, location, description string,
	scenary_spot_catagory_id int64) (int64, error) {

	scenary_spot := ScenarySpot{
		ScenarySpotCatagoryId: scenary_spot_catagory_id,
		EditTime:              time.Now().String(),
		Name:                  name,
		ShortName:             short_name,
		Location:              location,
		Description:           description}

	o := orm.NewOrm()
	id, err := o.Insert(&scenary_spot)

	return id, err
}

func UpdateScenarySpot(scenary_spot *ScenarySpot) int64 {
	o := orm.NewOrm()
	status, err := o.Update(scenary_spot)
	if err == nil {
		return status
	}
	return -1
}

func DeleveScenarySpot(scenary_spot_id int64) (num int64) {
	o := orm.NewOrm()
	num, _ = o.Delete(&ScenarySpot{Id: scenary_spot_id})
	return
}

func GetScenarySpot(scenary_spot_id int64) ScenarySpot {
	scenary_spot := ScenarySpot{Id: scenary_spot_id}
	o := orm.NewOrm()
	o.Read(&scenary_spot)

	return scenary_spot
}

func GetScenary(user_id string) map[string]interface{} {
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

func GetScenarySpots(number, page_size, page_index int64) map[string]interface{} {
	scenary_spots := make(map[string]interface{})

	o := orm.NewOrm()

	scenaries := make([]ScenarySpot, 10)

	number, err := o.Raw("select * from scenary_spot ").QueryRows(&scenaries)

	if err == nil {
		scenary_spots["scenaries"] = scenaries
	}

	scenary_spots["number"] = number
	scenary_spots["page_size"] = page_size
	scenary_spots["page_index"] = page_index

	return scenary_spots
}
