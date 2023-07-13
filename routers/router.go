// @APIVersion 1.0.0
// @Title Go Car Rental API
// @Description Static autogenerated API for Go Car Rental using Beego Swagger
// @Contact wojciech.bernatek@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
// @Name Go Car Rental
// @URL http://localhost:8080/
// @Host localhost:8080
// @Schemes HTTP, HTTPS
package routers

import (
	"beego-car-rental/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
		beego.NSNamespace("/auth",
			beego.NSInclude(
				&controllers.AuthController{},
			),
		),
		beego.NSNamespace("/car",
			beego.NSInclude(
				&controllers.CarController{},
			),
		),
		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),
		beego.NSNamespace("/privilege",
			beego.NSInclude(
				&controllers.PrivilegeController{},
			),
		),
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
