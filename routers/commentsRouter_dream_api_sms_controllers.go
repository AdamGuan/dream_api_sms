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
			`/register`,
			[]string{"post"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"],
		beego.ControllerComments{
			"ResetPwdGetSms",
			`/resetpwd/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"],
		beego.ControllerComments{
			"ResetPwd",
			`/resetpwd`,
			[]string{"put"},
			nil})

	beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"] = append(beego.GlobalControllerRouter["dream_api_sms/controllers:SmsController"],
		beego.ControllerComments{
			"CheckUserAndPwd",
			`/login/:mobilePhoneNumber`,
			[]string{"get"},
			nil})

}
