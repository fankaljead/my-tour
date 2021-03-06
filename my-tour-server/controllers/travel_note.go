package controllers

import (
	"fmt"
	"strconv"

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
	// user_id := int64(1)
	user_id := GetCurrentSessionUserIdInt64()

	title := u.GetString("title")
	content := u.GetString("content")
	routine_id, _ := u.GetInt64("routine_id")
	longitude := u.GetString("longitude")
	latitude := u.GetString("latitude")
	cover := u.GetString("cover")
	place := u.GetString("place")
	publishTime := u.GetString("publishTime")
	image_ids := u.GetString("image_ids")

	id, _ := models.AddTravelNote(title, content, cover, place, user_id, routine_id, longitude, latitude, image_ids, publishTime)
	u.Data["json"] = id

	u.ServeJSON()
}

// @Title UpdateTravelNote
// @Description  add travel note
// @Param	title path 	string	true
// @Param	content path 	string	true
// @Param	cover path 	string	true
// @Param	publishTime path 	string	true
// @Param	travel_note_id path 	int64	true
// @Success 200 {int} models.TravelNote.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TravelNoteController) UpdateTravelNote() {
	title := u.GetString("title")
	travel_note_id, _ := u.GetInt64("travel_note_id")
	content := u.GetString("content")
	cover := u.GetString("cover")
	place := u.GetString("place")
	publishTime := u.GetString("publishTime")

	// travel_note, err := models.GetTravelNote(travel_note_id)

	// if err == nil {
	// 	if (title )
	// }

	travel_note := models.TravelNote{
		Id:          travel_note_id,
		Title:       title,
		Content:     content,
		Cover:       cover,
		Place:       place,
		PublishTime: publishTime,
	}

	u.Data["json"] = models.UpdateTravelNote(&travel_note)
	u.ServeJSON()
}

// @Title GetTravelNoteById
// @Description get travel note by travel note id
// @Param	travel_note_id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} models.TravelNote
// @Failure 403 :uid is empty
// @router /:travel_note_id [get]
func (u *TravelNoteController) Get() {
	travel_note_id, _ := u.GetInt64(":travel_note_id")
	u.Data["json"], _ = models.GetTravelNote(travel_note_id)
	u.ServeJSON()
}

// @Title GetTravelNotesByRoutineId
// @Description get travel notes by travel routine id
// @Param	travel_routine_id		path 	int64	true		"The key for staticblock"
// @Success 200 {object} []models.TravelNote
// @Failure 403 :uid is empty
// @router /getTravelNotesInfoByRoutineId/:routine_id [get]
func (u *TravelNoteController) GetTravelNotesByRoutineId() {
	routine_id, _ := u.GetInt64(":routine_id")
	u.Data["json"] = models.GetTravelNoteByRoutineId(strconv.FormatInt(routine_id, 10))
	u.ServeJSON()
}

// @Title DeleteTravelNote
// @Description  delete travel note by id
// @Param	travel_note_id path 	int64	true		"The travel note you want to delete"
// @Success 200 {int} models.TravelNote.Id
// @Failure 403 body is empty
// @router / [delete]
func (u *TravelNoteController) Delete() {
	// user_id := 1
	travel_note_id, _ := u.GetInt64("travel_note_id")
	num := models.DeleveTravelNote(travel_note_id)

	u.Data["json"] = num

	u.ServeJSON()
}

// @Title GetAllTravelNoteInfoByUserId
// @Description get travel all note by user id
// @Success 200 {object} models.TravelNote
// @Failure 403 travel note info is empty
// @router /getAllTravelNoteInfo [get]
func (u *TravelNoteController) GetAllTravelNoteInfo() {
	user_id := GetCurrentSessionUserIdString()

	u.Data["json"] = models.GetAllUserTravelNoteInfo(user_id)
	u.ServeJSON()
}

// @Title GetAllSubscribeTravelNoteInfoByUserId
// @Description get travel all subscribe note by user id
// @Success 200 {object} models.TravelNote
// @Failure 403 travel note info is empty
// @router /getAllSubscribeTravelNoteInfo [get]
func (u *TravelNoteController) GetAllSubscribeTravelNoteInfo() {
	user_id := GetCurrentSessionUserIdString()
	u.Data["json"] = models.GetAllSubTravelNote(user_id)
	u.ServeJSON()
}

// @Title GetAllUserTravelNoteInfoByUserId
// @Description get travel all subscribe note by user id
// @Success 200 {object} models.TravelNote
// @Failure 403 travel note info is empty
// @router /getAllUserTravelNoteInfo [get]
func (u *TravelNoteController) GetAllUserTravelNoteInfo() {
	u.Data["json"] = models.GetAllUserTravelNote()
	u.ServeJSON()
}

// @Title SetTravelNoteDraft
// @Description set travel note as draft
// @Param	travel_note_id		 	int64	true		"body for user content"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /setTravelNoteDraft [post]
func (u *TravelNoteController) SetTravelNoteDraft() {

	travel_note_id, _ := u.GetInt64("travel_note_id")

	u.Data["json"] = models.SetTravelNoteDraft(travel_note_id)

	u.ServeJSON()
}

// @Title PublishTravelNoteDraft
// @Description publish travel note
// @Param	travel_note_id		 	int64	true		"body for user content"
// @Success 200 {int}
// @Failure 403 body is empty
// @router /publishTravelNote [post]
func (u *TravelNoteController) PublishTravelNoteDraft() {

	travel_note_id, _ := u.GetInt64("travel_note_id")

	u.Data["json"] = models.PublishTravelNoteDraft(travel_note_id)

	u.ServeJSON()
}

// @Title AddTravelRoutine
// @Description  add travel note ??????????????????
// @Param	title path 	string	true
// @Param	content path 	string	true
// @Param	cover path 	string	true
// @Param	publishTime path 	string	true
// @Success 200 {int} models.TravelNote.Id
// @Failure 403 body is empty
// @router /add_travel_routine [post]
func (u *TravelNoteController) AddTravelRoutine() {
	// user_id := int64(1)
	user_id := GetCurrentSessionUserIdInt64()

	name := u.GetString("title")
	description := u.GetString("description")

	id, _ := models.AddTravelRoutine(name, description, user_id)
	u.Data["json"] = id

	u.ServeJSON()
}

// @Title GetTravelRoutines
// @Description get travel all routines by user id
// @Success 200 {object} models.TravelRoutine
// @Failure 403 travel routine is empty
// @router /get_travel_routines [get]
func (u *TravelNoteController) GetAllTravelRoutines() {
	user_id := GetCurrentSessionUserIdInt64()

	u.Data["json"] = models.GetTravelRoutines(1, 100, user_id)

	u.ServeJSON()
}

// @Title DeleteTravelRoutine
// @Description  delete travel routine by id
// @Param	travel_routine_id path 	int64	true		"The travel routine you want to delete"
// @Success 200 {int} models.TravelNoteRoutine.Id
// @Failure 403 body is empty
// @router /delete_travel_routine [delete]
func (u *TravelNoteController) DeleteTravelNoteRoutine() {
	// user_id := 1
	travel_routine_id, _ := u.GetInt64("travel_routine_id")
	fmt.Printf("======= delete id = %d =====", travel_routine_id)
	num := models.DeleteTravelRoutine(travel_routine_id)

	u.Data["json"] = num

	u.ServeJSON()
}
