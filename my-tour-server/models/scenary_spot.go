package models

import (
	"strconv"
	"time"

	"github.com/astaxie/beego/orm"
)

type ScenarySpot struct {
	Id                    int64   `json:"id"`
	Name                  string  `json:"name"`
	ShortName             string  `json:"shortName"`
	Background            string  `json:"background"`
	Location              string  `json:"location"`
	Description           string  `json:"description"`
	EditTime              string  `json:"editTime"`
	Score                 float64 `json:"score"`
	ScenarySpotCatagoryId int64   `json:"scenary_spot_catagory_id"`
}

type ScenarySpotCatagory struct {
	Id           int64
	Name         string
	Description  string
	CreatTime    time.Time `orm:"auto_now_add;type(datetime)"`
	CreateUserId int64
}

func init() {
	orm.RegisterModel(new(ScenarySpot), new(ScenarySpotCatagory))
	AddATable("scenary_spot")
	AddATable("scenary_spot_catagory")
}

func AddScenarySpot(name, short_name, location, description, background string,
	scenary_spot_catagory_id int64) (int64, error) {

	scenary_spot := ScenarySpot{
		ScenarySpotCatagoryId: scenary_spot_catagory_id,
		EditTime:              time.Now().String(),
		Name:                  name,
		ShortName:             short_name,
		Location:              location,
		Background:            background,
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

func GetScenarySpots(page_size, page_index int64) map[string]interface{} {
	scenary_spots := make(map[string]interface{})

	o := orm.NewOrm()

	scenaries := make([]ScenarySpot, 10)

	start := page_size * (page_index - 1)
	end := page_size * page_index

	number, err := o.Raw("select * from scenary_spot limit " +
		strconv.FormatInt(start, 10) +
		"," + strconv.FormatInt(end, 10)).QueryRows(&scenaries)

	if err == nil {
		scenary_spots["scenaries"] = scenaries
	}

	scenary_spots["number"] = number
	scenary_spots["page_size"] = page_size
	scenary_spots["page_index"] = page_index

	return scenary_spots
}

func GetScenarySpotsShortInfo(page_size, page_index int64) map[string]interface{} {
	scenary_spots := make(map[string]interface{})

	o := orm.NewOrm()

	// scenaries := make([]ScenarySpot, 10)

	var scenaries []struct {
		Id   int64  `json:"id"`
		Name string `json:"name"`
	}

	start := page_size * (page_index - 1)
	end := page_size * page_index

	number, err := o.Raw("select id,name from scenary_spot limit " +
		strconv.FormatInt(start, 10) +
		"," + strconv.FormatInt(end, 10)).QueryRows(&scenaries)

	if err == nil {
		scenary_spots["scenaries"] = scenaries
	}

	scenary_spots["number"] = number
	scenary_spots["page_size"] = page_size
	scenary_spots["page_index"] = page_index

	return scenary_spots
}

func AddScenarySpotCatagory(name, description string, create_user_id int64) (int64, error) {
	scenary_spot_catagory := ScenarySpotCatagory{
		Name:         name,
		Description:  description,
		CreateUserId: create_user_id,
	}

	o := orm.NewOrm()
	id, err := o.Insert(&scenary_spot_catagory)

	return id, err
}

func DeleveScenarySpotCatagory(scenary_spot_catagory_id int64) (num int64) {
	o := orm.NewOrm()
	num, _ = o.Delete(&ScenarySpotCatagory{Id: scenary_spot_catagory_id})
	return
}

func UpdateScenarySpotCatagory(scenary_spot_catagory *ScenarySpotCatagory) int64 {
	o := orm.NewOrm()
	status, err := o.Update(scenary_spot_catagory)
	if err == nil {
		return status
	}
	return -1
}

func GetScenarySpotsCatagory(page_size, page_index int64) map[string]interface{} {
	scenary_spots_catagory := make(map[string]interface{})

	o := orm.NewOrm()

	scenary_catagories := make([]ScenarySpotCatagory, 10)

	start := page_size * (page_index - 1)
	end := page_size * page_index

	number, err := o.Raw("select * from scenary_spot_catagory limit " +
		strconv.FormatInt(start, 10) +
		"," + strconv.FormatInt(end, 10)).QueryRows(&scenary_catagories)

	if err == nil {
		scenary_spots_catagory["scenaries"] = scenary_catagories
	}

	scenary_spots_catagory["number"] = number
	scenary_spots_catagory["page_size"] = page_size
	scenary_spots_catagory["page_index"] = page_index

	return scenary_spots_catagory
}
