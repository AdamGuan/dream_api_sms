package controllers

import (
	"dream_api_sms/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms/helper"
	//"fmt"
	//"strings"
)

//短信
type SmsController struct {
	beego.Controller
}

// @Title 发送一条短信验证码(注册时)
// @Description 发送一条短信验证码(注册时)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
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
	pkg := u.Ctx.Request.Header.Get("Pkg")
	sign := u.Ctx.Request.Header.Get("Sign")
	datas["responseNo"] = -6
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		datas["responseNo"] = -7
	}else{	
		if result := helper.CheckSign(sign, pkg); result {
			datas["responseNo"] = 0
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 {
		datas["responseNo"] = -1
		res2 := userObj.CheckUserNameValid(mobilePhoneNumber)
		if res2 == 0{
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0{
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
				if len(res) == 0{
					datas["responseNo"] = 0
				}
			}
		}else{
			datas["responseNo"] = res2
		}
	}

	//return
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 注册
// @Description 注册
// @Param	mobilePhoneNumber	form	string	true	"手机号码"
// @Param	num			form	string	true	"验证码"
// @Param	pwd			form	string	true	"密码"
// @Param	sign			header	string	true	"签名"
// @Param	pkg			header	string	true	"包名"
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register [post]
func (u *SmsController) Register() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	num := u.Ctx.Request.FormValue("num")
	pwd := u.Ctx.Request.FormValue("pwd")
	//check sign
	pkg := u.Ctx.Request.Header.Get("Pkg")
	sign := u.Ctx.Request.Header.Get("Sign")
	datas["responseNo"] = -6
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		datas["responseNo"] = -7
	}else{	
		if result := helper.CheckSign(sign, pkg); result {
			datas["responseNo"] = 0
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 && len(num) > 0 && len(pwd) > 0 {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0{
			res := smsObj.ValidMsm(num,mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
			if len(res) == 0{
				res2 := userObj.AddUser(mobilePhoneNumber,pwd)
				datas["responseNo"] = res2
			}
		}
	}
	//return
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 发送一条短信验证码(重置密码时)
// @Description 发送一条短信验证码(重置密码时)
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd/:mobilePhoneNumber [get]
func (u *SmsController) ResetPwdGetSms() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	pkg := u.Ctx.Request.Header.Get("Pkg")
	sign := u.Ctx.Request.Header.Get("Sign")
	datas["responseNo"] = -6
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		datas["responseNo"] = -7
	}else{	
		if result := helper.CheckSign(sign, pkg); result {
			datas["responseNo"] = 0
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 {
		datas["responseNo"] = -1
		res := userObj.CheckUserNameExists(mobilePhoneNumber)
		if res {
			pkgConfig := pkgObj.GetPkgConfig(pkg)
			if len(pkgConfig) > 0 {
				res := smsObj.GetMsm(mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
				if len(res) == 0{
					datas["responseNo"] = 0
				}
			}
		}else{
			datas["responseNo"] = -4
		}
	}

	//return
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 重置密码
// @Description 重置密码
// @Param	mobilePhoneNumber	form	string	true	手机号码
// @Param	num			form	string	true	验证码
// @Param	pwd			form	string	true	密码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd [put]
func (u *SmsController) ResetPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var smsObj *models.MSms
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	num := u.Ctx.Request.FormValue("num")
	pwd := u.Ctx.Request.FormValue("pwd")
	//check sign
	pkg := u.Ctx.Request.Header.Get("Pkg")
	sign := u.Ctx.Request.Header.Get("Sign")
	datas["responseNo"] = -6
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		datas["responseNo"] = -7
	}else{	
		if result := helper.CheckSign(sign, pkg); result {
			datas["responseNo"] = 0
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 && len(num) > 0 && len(pwd) > 0 {
		datas["responseNo"] = -1
		pkgConfig := pkgObj.GetPkgConfig(pkg)
		if len(pkgConfig) > 0{
			res := smsObj.ValidMsm(num,mobilePhoneNumber,pkgConfig["F_app_id"],pkgConfig["F_app_key"])
			if len(res) == 0{
				res2 := userObj.ModifyUserPwd(mobilePhoneNumber,pwd)
				datas["responseNo"] = res2
			}
		}
	}
	//return
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 登录验证
// @Description 登录验证
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	pwd			query	string	true	密码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /login/:mobilePhoneNumber [get]
func (u *SmsController) CheckUserAndPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")	
	//check sign
	pkg := u.Ctx.Request.Header.Get("Pkg")
	sign := u.Ctx.Request.Header.Get("Sign")
	datas["responseNo"] = -6
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		datas["responseNo"] = -7
	}else{	
		if result := helper.CheckSign(sign, pkg); result {
			datas["responseNo"] = 0
		}
	}
	//检查参数
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 && len(pwd) > 0 {
		datas["responseNo"] = -1
		res := userObj.CheckUserAndPwd(mobilePhoneNumber,pwd)
		if res{
			datas["responseNo"] = 0
		}else{
			datas["responseNo"] = -5
		}
	}
	//return
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}