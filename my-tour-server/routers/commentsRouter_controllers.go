package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:FansController"],
        beego.ControllerComments{
            Method: "FollowUser",
            Router: `/followUser`,
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
            Router: `/unfollowUser`,
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

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ImageController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ImageController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ImageController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ImageController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ImageController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ImageController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
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

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "AddScenarySpot",
            Router: `/addScenarySpot`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "UpdateScenarySpot",
            Router: `/addScenarySpot`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "AddScenarySpotCatagory",
            Router: `/addScenarySpotCatagory`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "DeleteScenarySpot",
            Router: `/deleteScenarySpot`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "DeleteScenarySpotCatagory",
            Router: `/deleteScenarySpotCatagory`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "GetScenarySpot",
            Router: `/getScenarySpot`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "GetScenarySpotCatagory",
            Router: `/getScenarySpotCatagory`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "GetScenarySpots",
            Router: `/getScenarySpots`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:ScenarySpotController"],
        beego.ControllerComments{
            Method: "GetScenarySpotsShorInfo",
            Router: `/getScenarySpotsShortInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TableController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TableController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "UpdateTravelNote",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:travel_note_id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "AddTravelRoutine",
            Router: `/add_travel_routine`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "DeleteTravelNoteRoutine",
            Router: `/delete_travel_routine`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "GetAllTravelNoteInfo",
            Router: `/getAllTravelNoteInfo`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "GetAllTravelRoutines",
            Router: `/get_travel_routines`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "PublishTravelNoteDraft",
            Router: `/publishTravelNote`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"] = append(beego.GlobalControllerRouter["gitlab.com/fankaljead/my-tour/my-tour-server/controllers:TravelNoteController"],
        beego.ControllerComments{
            Method: "SetTravelNoteDraft",
            Router: `/setTravelNoteDraft`,
            AllowHTTPMethods: []string{"post"},
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
