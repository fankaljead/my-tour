package models

import "github.com/astaxie/beego/orm"

type ScenarySpotCommet struct {
	Id            int64   `json:"id"`
	Title         string  `json:"title"`
	Content       string  `json:"content"`
	UserId        int64   `json:"user_id"`
	ScenarySpotId int64   `json:"scenary_spot_id"`
	StarNumber    int     `json:"star_number"`
	Score         float64 `json:"score"`
}

func init() {
	orm.RegisterModel(new(ScenarySpotCommet))
	AddATable("scenary_spot_comment")
}
