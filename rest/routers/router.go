package routers

import (
	"github.com/awethome/server/rest/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",

		beego.NSNamespace("/action",
			beego.NSInclude(
				&controllers.ActionController{},
			),
		),

		beego.NSNamespace("/realm",
			beego.NSInclude(
				&controllers.RealmController{},
			),
		),

		beego.NSNamespace("/role",
			beego.NSInclude(
				&controllers.RoleController{},
			),
		),

		beego.NSNamespace("/user",
			beego.NSInclude(
				&controllers.UserController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
