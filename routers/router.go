// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"helloapp/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

/*
*
beego.NewNamespace 是创建命名空间

beego.NSNamespace 是希望namespace内部嵌套namespace时使用，NSNamespace 来注入一个子namespace。

beego.NSInclude 该方法是注解路由的配套方法，只对注解路由生效。
*/
func init() {
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
		beego.NSNamespace("/order",
			beego.NSInclude(
				&controllers.OrderController{},
			),
		),
		beego.NSNamespace("/default",
			beego.NSInclude(
				&controllers.MainController{},
			),
		),
	)
	beego.AddNamespace(ns)

	beego.Router("/PhotoDetail", &controllers.MainController{}, "*:PhotoDetail")
	beego.Router("/", &controllers.OrderController{}, "*:Test")
	beego.Router("/v2/test2", &controllers.OrderController{}, "*:Test2")
	beego.Include(&controllers.OrderController{})
}
