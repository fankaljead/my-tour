package utils

import (
	"fmt"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/controllers"
)

// get current session user id
func GetSessionUserId(c *beego.Controller) string {

	r := c.Ctx.Request
	w := c.Ctx.ResponseWriter
	sess, _ := controllers.GlobalSessions.SessionStart(w, r)
	defer sess.SessionRelease(w)
	user_id := fmt.Sprintf("%v", sess.Get("user_id"))

	return user_id
}
