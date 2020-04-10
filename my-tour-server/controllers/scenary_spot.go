package controllers

import (
	"github.com/astaxie/beego"
	"gitlab.com/fankaljead/my-tour/my-tour-server/models"
)

type ScenarySpotController struct {
	beego.Controller
}

// @Title CreateAScenarySpot
// @Description create object
// @Param	name		sting	true		"The spot name"
// @Param	short_name		sting	true		"The spot short_name"
// @Param	location		sting	true		"The spot location"
// @Param	description		sting	true		"The spot description"
// @Param	scenary_spot_catagory_id		int64	true		"The spot scenary_spot_catagory_id"
// @Success 200 {string} models.ScenarySpot.Id
// @Failure 403 body is empty
// @router /addScenarySpot [post]
func (u *ScenarySpotController) AddScenarySpot() {
	name := u.GetString("name")
	short_name := u.GetString("short_name")
	location := u.GetString("location")
	description := u.GetString("description")
	background := u.GetString("background")
	scenary_spot_catagory_id, _ := u.GetInt64("scenary_spot_catagory_id")

	id, _ := models.AddScenarySpot(name, short_name, location, description, background, scenary_spot_catagory_id)

	u.Data["json"] = id

	u.ServeJSON()
}

// @Title UpdateScenarySpot
// @Description update scenary spot
// @Param	scenary_spot_id int64	true		"The spot scenary_spot_id"
// @Param	name		sting	true		"The spot name"
// @Param	short_name		sting	true		"The spot short_name"
// @Param	location		sting	true		"The spot location"
// @Param	description		sting	true		"The spot description"
// @Param	scenary_spot_catagory_id		int64	true		"The spot scenary_spot_catagory_id"
// @Success 200 {string} models.ScenarySpot.Id
// @Failure 403 body is empty
// @router /addScenarySpot [post]
func (u *ScenarySpotController) UpdateScenarySpot() {
	name := u.GetString("name")
	short_name := u.GetString("short_name")
	location := u.GetString("location")
	description := u.GetString("description")
	background := u.GetString("background")
	scenary_spot_catagory_id, _ := u.GetInt64("scenary_spot_catagory_id")
	scenary_spot_id, _ := u.GetInt64("scenary_spot_id")

	scenary_spot := models.ScenarySpot{
		Id:                    scenary_spot_id,
		Name:                  name,
		ShortName:             short_name,
		Location:              location,
		Description:           description,
		Background:            background,
		ScenarySpotCatagoryId: scenary_spot_catagory_id,
	}

	status := models.UpdateScenarySpot(&scenary_spot)

	u.Data["json"] = status

	u.ServeJSON()
}

// @Title GetScenarySpot
// @Description get scenary spot
// @Param	scenary_spot_id		int64	true		"The scenary spot id "
// @Success 200 {object} models.ScenarySpot
// @Failure 403 body is empty
// @router /getScenarySpot [get]
func (u *ScenarySpotController) GetScenarySpot() {

	scenary_spot_id, _ := u.GetInt64("scenary_spot_id")

	u.Data["json"] = models.GetScenarySpot(scenary_spot_id)

	u.ServeJSON()
}

// @Title GetScenarySpots
// @Description get scenary spots
// @Param	pase_size		int64	true		"The page size "
// @Param	page_index		int64	true		"The page index"
// @Success 200 {string} models.ScenarySpot
// @Failure 403 body is empty
// @router /getScenarySpots [get]
func (u *ScenarySpotController) GetScenarySpots() {

	page_size, _ := u.GetInt64("page_size")
	page_index, _ := u.GetInt64("page_index")

	u.Data["json"] = models.GetScenarySpots(page_size, page_index)

	u.ServeJSON()
}

// @Title GetScenarySpotsShortInfo
// @Description get scenary spots short info
// @Param	pase_size		int64	true		"The page size "
// @Param	page_index		int64	true		"The page index"
// @Success 200 {objec} models.ScenarySpot
// @Failure 403 body is empty
// @router /getScenarySpotsShortInfo [get]
func (u *ScenarySpotController) GetScenarySpotsShorInfo() {

	page_size, _ := u.GetInt64("page_size")
	page_index, _ := u.GetInt64("page_index")

	u.Data["json"] = models.GetScenarySpotsShortInfo(page_size, page_index)

	u.ServeJSON()
}

// @Title DeleteScenarySpot
// @Description  delete scenary spot
// @Param	scenary_spot_id path 	int64	true		"The scenary spot you want to delete"
// @Success 200 {int} models.ScenarySpot.Id
// @Failure 403 body is empty
// @router /deleteScenarySpot [delete]
func (u *ScenarySpotController) DeleteScenarySpot() {
	// user_id := 1
	scenary_spot_id, _ := u.GetInt64("scenary_spot_id")

	num := models.DeleveScenarySpot(scenary_spot_id)

	u.Data["json"] = num

	u.ServeJSON()
}

// @Title CreateAScenarySpotCatagory
// @Description create object
// @Param	name		sting	true		"The spot name"
// @Param	description		sting	true		"The spot description"
// @Param	create_user_id		int64	true		"The spot create_user_id"
// @Success 200 {string} models.ScenarySpotCatagory.Id
// @Failure 403 body is empty
// @router /addScenarySpotCatagory [post]
func (u *ScenarySpotController) AddScenarySpotCatagory() {
	name := u.GetString("name")
	description := u.GetString("description")
	create_user_id, _ := u.GetInt64("create_user_id")

	id, _ := models.AddScenarySpotCatagory(name, description, create_user_id)

	u.Data["json"] = id

	u.ServeJSON()
}

// @Title GetScenarySpotCatagory
// @Description get scenary spot catagory
// @Param	pase_size		int64	true		"The page size "
// @Param	page_index		int64	true		"The page index"
// @Success 200 {string} models.ScenarySpotCatagory
// @Failure 403 body is empty
// @router /getScenarySpotCatagory [get]
func (u *ScenarySpotController) GetScenarySpotCatagory() {

	page_size, _ := u.GetInt64("page_size")
	page_index, _ := u.GetInt64("page_index")

	u.Data["json"] = models.GetScenarySpotsCatagory(page_size, page_index)

	u.ServeJSON()
}

// @Title DeleteScenarySpotCatagory
// @Description  delete scenary spot catagory
// @Param	scenary_spot_catagory_id path 	int64	true		"The scenary spot catagory you want to delete"
// @Success 200 {int} models.ScenarySpotCatagory.Id
// @Failure 403 body is empty
// @router /deleteScenarySpotCatagory [delete]
func (u *ScenarySpotController) DeleteScenarySpotCatagory() {

	scenary_spot_catagory_id, _ := u.GetInt64("scenary_spot_catagory_id")

	num := models.DeleveScenarySpotCatagory(scenary_spot_catagory_id)

	u.Data["json"] = num

	u.ServeJSON()
}
