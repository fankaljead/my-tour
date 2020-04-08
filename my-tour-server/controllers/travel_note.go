package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

type TravelNoteController struct {
	beego.Controller
}

// @Title AddTravelNote
// @Description  add travel note
// @Param	title path 	string	true
// @Param	content path 	string	true
// @Param	cover path 	string	true
// @Param	publishTime path 	string	true
// @Success 200 {int} models.TravelNote.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TravelNoteController) Post() {
	user_id := int64(1)

	title := u.GetString("title")
	content := u.GetString("content")
	cover := u.GetString("cover")
	place := u.GetString("place")
	publishTime := u.GetString("publishTime")

	id, _ := models.AddTravelNote(title, content, cover, place, user_id, publishTime)
	u.Data["json"] = id

	u.ServeJSON()
}
