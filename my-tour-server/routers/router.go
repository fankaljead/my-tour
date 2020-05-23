// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"gitlab.com/fankaljead/my-tour/my-tour-server/controllers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
)

func init() {
	beego.Debug("Filters init...")
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Access-Control-Allow-Origin"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin"},
		AllowCredentials: true,
	}))
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/object",
			beego.NSInclude(
				&controllers.ObjectController{},
			),
		),
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/user_info",
			beego.NSInclude(
				&controllers.UserInfoController{},
			),
		),
		beego.NSNamespace("/hobby",
			beego.NSInclude(
				&controllers.HobbyController{},
			),
		),
		beego.NSNamespace("/fan",
			beego.NSInclude(
				&controllers.FansController{},
			),
		),
		beego.NSNamespace("/travel_note",
			beego.NSInclude(
				&controllers.TravelNoteController{},
			),
		),
		beego.NSNamespace("/scenary_spot",
			beego.NSInclude(
				&controllers.ScenarySpotController{},
			),
		),
		beego.NSNamespace("/table",
			beego.NSInclude(
				&controllers.TableController{},
			),
		),
		beego.NSNamespace("/image",
			beego.NSInclude(
				&controllers.ImageController{},
			),
		),
		beego.NSNamespace("/comment",
			beego.NSInclude(
				&controllers.TravelNoteCommentController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
