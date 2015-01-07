package controllers

import (
	"dream_api_sms/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms/helper"
	//"fmt"
)

//短信
type SmsController struct {
	beego.Controller
}

// @Title 发送一条短信验证码(注册时)
// @Description 发送一条短信验证码(注册时)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	false	签名(保留)
// @Param	pkg			header	string	false	包名(保留)
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register/:mobilePhoneNumber [get]
func (u *SmsController) RegisterGetSms() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = 0
	/*
	if result := helper.CheckSign(u.Ctx.Request.Form, key); result {
		datas["responseNo"] = 0
	}
	*/

	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 {
		datas["responseNo"] = -1
		res2 := userObj.CheckUserNameValid(mobilePhoneNumber)
		if res2 == 0{
			res := smsObj.GetMsm(mobilePhoneNumber)
			if len(res) == 0{
				datas["responseNo"] = 0
			}
		}else{
			datas["responseNo"] = res2
		}
	}

	//return
	if datas["responseNo"] == -1 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 注册
// @Description 注册
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	num			path	string	true	验证码
// @Param	pwd			query	string	true	密码
// @Param	sign			header	string	false	签名(保留)
// @Param	pkg			header	string	false	包名(保留)
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register/:mobilePhoneNumber/:num [get]
func (u *SmsController) Register() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	num := u.Ctx.Input.Param(":num")
	pwd := u.GetString("pwd")
	//check sign
	datas["responseNo"] = 0
	/*
	if result := helper.CheckSign(u.Ctx.Request.Form, key); result {
		datas["responseNo"] = 0
	}
	*/

	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 && len(num) > 0 && len(pwd) > 0 {
		datas["responseNo"] = -1
		res := smsObj.ValidMsm(num,mobilePhoneNumber)
		if len(res) == 0{
			res2 := userObj.AddUser(mobilePhoneNumber,pwd)
			datas["responseNo"] = res2
		}
	}else{
		datas["responseNo"] = -1
	}

	//return
	if datas["responseNo"] == -1 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}