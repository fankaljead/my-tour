package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

type FansController struct {
	beego.Controller
}

// @Title FollowUser
// @Description  follow user
// @Param	follower_id path 	int64	true		"The user id you want to follow"
// @Success 200 {int} models.Fans.Id
// @Failure 403 body is empty
// @router /followUser [post]
func (u *FansController) FollowUser() {
	user_id := int64(1)

	follower_id, _ := u.GetInt64("follower_id")

	id, _ := models.FollowUser(user_id, follower_id)
	u.Data["json"] = id

	u.ServeJSON()
}

// @Title UnfollowUser
// @Description unfollow user
// @Param	follower_id path 	int64	true		"The user id you want to follow"
// @Success 200 {int} models.Fans.Id
// @Failure 403 body is empty
// @router /unfollowUser [post]
func (u *FansController) UnfollowUser() {
	user_id := int64(1)

	follower_id, _ := u.GetInt64("follower_id")

	num := models.UnfollowUser(user_id, follower_id)
	u.Data["json"] = num

	u.ServeJSON()
}

// @Title GetAllFollowers
// @Description get all user followers
// @Success 200 {map[string]interface{}}
// @router /getAllFollowers [get]
func (u *FansController) GetAllFollower() {
	// user_id := utils.GetSessionUserId(u)
	user_id := "1"
	u.Data["json"] = models.GetAllFollowers(user_id)
	u.ServeJSON()
}
