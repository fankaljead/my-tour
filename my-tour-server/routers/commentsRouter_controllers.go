package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"],
        beego.ControllerComments{
            Method: "FollowUser",
            Router: `/follow_user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"],
        beego.ControllerComments{
            Method: "GetAllFollower",
            Router: `/getAllFollowers`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"],
        beego.ControllerComments{
            Method: "UnfollowUser",
            Router: `/unfollow_user`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:HobbyController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:HobbyController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:HobbyController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:HobbyController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:HobbyController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:HobbyController"],
        beego.ControllerComments{
            Method: "GetAllUserHobbies",
            Router: `/getAllUserHobbies`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:user_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "SetBackgroundImage",
            Router: `/set_background_image`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:UserInfoController"],
        beego.ControllerComments{
            Method: "SetIcon",
            Router: `/set_icon`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
