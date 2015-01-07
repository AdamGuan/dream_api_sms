// @APIVersion 1.0.0
// @Title 短信验证 API
package routers

import (
	"dream_api_sms/controllers"

	"github.com/astaxie/beego"
)

func init() {
	ns := beego.NewNamespace("/v1",
		beego.NSNamespace("/sms",
			beego.NSInclude(
				&controllers.SmsController{},
			),
		),
	)
	beego.AddNamespace(ns)
}
