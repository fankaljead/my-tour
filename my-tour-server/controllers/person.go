package controllers

import (
	"encoding/json"

	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

type PersonController struct {
	beego.Controller
}

// @Title CreatePerson
// @Description create persons
// @Param	body		body 	models.Person	true		"body for person content"
// @Success 200 {int} models.Person.Id
// @Failure 403 body is empty
// @router / [post]
func (u *PersonController) Post() {
	var person models.Person
	// json.Unmarshal(u.Ctx.Input.RequestBody, &person)
	u.ParseForm(&person)
	uid := models.AddPerson(person)
	u.Data["json"] = map[string]string{"uid": uid}

	u.ServeJSON()
}

// @Title Get
// @Description get user by uid
// @Success 200 {object} models.Person
// @Failure 403 :uid is empty
// @router / [get]
// func (u *PersonController) Get() {
// 	u.Data["data"] = "person controller"
// 	u.ServeJSON()
// }

// @Title Update
// @Description update the person
// @Param id path string true
// @Success 200 {object} models.Person
// @Failure 403 :uid is not int
// @router /:uid [put]
func (p *PersonController) Put() {
	id := p.GetString(":id")
	if uid != "" {
		var person models.Person
		json.Unmarshal(p.Ctx.Input.RequestBody, &person)
	}
}
