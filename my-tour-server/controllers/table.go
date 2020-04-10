package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

type TableController struct {
	beego.Controller
}

// @Title CreateTable
// @Description  create table
// @Success 200 {int}
// @Failure 403 body is empty
// @router / [post]
func (u *TableController) Post() {
	successNum, err := models.AutoCreateTabel()
	if err == nil {
		u.Data["json"] = successNum
	}

	u.ServeJSON()
}
