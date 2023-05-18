package routers

import (
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context/param"
)

func init() {

    beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:id`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:id`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:CarController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:id`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uuid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uuid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uuid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"] = append(beego.GlobalControllerRouter["beego-car-rental/controllers:UserController"],
        beego.ControllerComments{
            Method: "Register",
            Router: `/register`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
