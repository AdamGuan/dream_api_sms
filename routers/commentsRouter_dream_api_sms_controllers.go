package routers

import (
	"github.com/astaxie/beego"
)

func init() {
	
	beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"],
		beego.ControllerComments{
			"RegisterGetSms",
			`/register/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"],
		beego.ControllerComments{
			"Register",
			`/register/:mobilePhoneNumber/:num`,
			[]string{"get"},
			nil})

}
