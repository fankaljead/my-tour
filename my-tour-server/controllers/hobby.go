package controllers

import (
	"fmt"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

// Operations about UserInfo
type HobbyController struct {
	beego.Controller
}

// @Title AddUserHobby
// @Description  add user hobby
// @Param	title path 	string	true		"The hobby you want to add"
// @Success 200 {int} models.UserHobby.Id
// @Failure 403 body is empty
// @router / [post]
func (u *HobbyController) Post() {
	// user_id := 1
	title := u.GetString("title")
	user_id := int64(1)
	hobby_id := models.AddHobby(title)
	id := models.AddUserHobby(user_id, hobby_id)

	u.Data["json"] = id

	u.ServeJSON()
}

// @Title DeleteUserHobby
// @Description  delete user hobby
// @Param	user_hobby_id path 	int64	true		"The hobby you want to delete"
// @Success 200 {int} models.UserHobby.Id
// @Failure 403 body is empty
// @router / [delete]
func (u *HobbyController) Delete() {
	// user_id := 1
	user_hobby_id, _ := u.GetInt64("user_hobby_id")
	fmt.Printf("user hobby id string is : %s \n", u.GetString("user_hobby_id"))
	fmt.Printf("user hobby id is : %d \n", user_hobby_id)
	num, _ := models.DeleteUserHobby(user_hobby_id)

	u.Data["json"] = num

	u.ServeJSON()
}

// @Title GetAllUserHobbies
// @Description get all user hobbies
// @Success 200 {map[string]interface{}}
// @router /getAllUserHobbies [get]
func (u *HobbyController) GetAllUserHobbies() {
	// user_id := utils.GetSessionUserId(u)
	user_id := "1"
	u.Data["json"] = models.GetUserHobbies(user_id)
	u.ServeJSON()
}
