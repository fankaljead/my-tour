package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

type TravelNoteCommentController struct {
	beego.Controller
}

// @Title Add A Travel Note comment
// @Description  follow user
// @Param	content path 	string	true		"The content you want to comment"
// @Param	travel_note_id path 	int64	true		"The travel_note_id you want to comment"
// @Success 200 {int} models.TravelNoteComment.Id
// @Failure 403 body is empty
// @router / [post]
func (u *TravelNoteCommentController) Post() {
	user_id := GetCurrentSessionUserIdInt64()

	content := u.GetString("content")
	travel_note_id, _ := u.GetInt64("travel_note_id")

	id, _ := models.AddATravelNoteCommet(content, travel_note_id, user_id)
	u.Data["json"] = id

	u.ServeJSON()
}

// @Title Get Travel Note Comments
// @Description get user travel note comments
// @Success 200 {map[string]interface{}}
// @router / [get]
func (u *TravelNoteCommentController) Get() {

	travel_note_id, _ := u.GetInt64("travel_note_id")

	u.Data["json"] = models.GetTravelNoteComments(100, 1, travel_note_id)
	u.ServeJSON()
}

// @Title Delete TravelNoteComment
// @Description delete the user
// @Param	travel_note_comment_id		path 	string	true		"The travel_note_comment_id you want to delete"
// @Success 200 {int} 1 delete success!
// @Failure 403 0 travel_note_comment_id is empty
// @router /:travel_note_comment_id [delete]
func (u *TravelNoteCommentController) Delete() {
	// u.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	// u.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", "Origin, X-Requested-With, Content-Type, Accept")
	travel_note_comment_id, _ := u.GetInt64(":travel_note_comment_id")
	num := models.DeleveTravelNoteComment(travel_note_comment_id)
	u.Data["json"] = num
	u.ServeJSON()
}

// @Title UpdateComment
// @Description  update comment
// @Param	comment_id path 	int	true
// @Param	content path 	string	true
// @Success 200 {int} models.Comment.Id
// @Failure 403 body is empty
// @router /update [post]
func (u *TravelNoteCommentController) UpdateTravelNoteComment() {
	comment_id, _ := u.GetInt64("comment_id")
	content := u.GetString("content")
	u.Data["json"], _ = models.UpdateTravelNoteComment(comment_id, content)
	u.ServeJSON()
}
