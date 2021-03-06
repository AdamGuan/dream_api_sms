package controllers

import (
	"dream_api_sms/models"
	"github.com/astaxie/beego"
	"net/http"
	"dream_api_sms/helper"
	//"fmt"
	//"strings"
)

//用户
type UserController struct {
	beego.Controller
}

//json echo
func (u0 *UserController) jsonEcho(datas map[string]interface{},u *UserController) {
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

//sign check
func (u0 *UserController) checkSign(u *UserController)int {
	result := -6
	pkg := u.Ctx.Request.Header.Get("Pkg")
	sign := u.Ctx.Request.Header.Get("Sign")
	var pkgObj *models.MPkg
	if !pkgObj.CheckPkgExists(pkg){
		result = -7
	}else{	
		if re := helper.CheckSign(sign, pkg); re {
			result = 0
		}
	}
	return result
}

// @Title 注册
// @Description 注册
// @Param	mobilePhoneNumber	form	string	true	"手机号码"
// @Param	pwd			form	string	true	"密码"
// @Param	sign			header	string	true	"签名"
// @Param	pkg			header	string	true	"包名"
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /register [post]
func (u *UserController) Register() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		res2 := userObj.AddUser(mobilePhoneNumber,pwd)
		datas["responseNo"] = res2
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 重置密码
// @Description 重置密码
// @Param	mobilePhoneNumber	form	string	true	手机号码
// @Param	pwd			form	string	true	密码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /resetpwd [put]
func (u *UserController) ResetPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		res2 := userObj.ModifyUserPwd(mobilePhoneNumber,pwd)
		datas["responseNo"] = res2
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
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
func (u *UserController) CheckUserAndPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	pwd := u.Ctx.Request.FormValue("pwd")	
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(pwd) {
		datas["responseNo"] = -1
		if !userObj.CheckUserNameExists(mobilePhoneNumber){
			datas["responseNo"] = -4
		}else{
			res := userObj.CheckUserAndPwd(mobilePhoneNumber,pwd)
			if res{
				datas["responseNo"] = 0
			}else{
				datas["responseNo"] = -9
			}
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -5
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 找回密码
// @Description 找回密码
// @Param	mobilePhoneNumber	path	string	true	"手机号码"
// @Param	sign			header	string	true	"签名"
// @Param	pkg			header	string	true	"包名"
// @Success	200 {object} models.MFindPwdResp
// @Failure 401 无权访问
// @router /pwd/:mobilePhoneNumber [get]
func (u *UserController) FindPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		datas["responseNo"] = -1
		if userObj.CheckUserNameExists(mobilePhoneNumber){
			res := userObj.GetUserPwd(mobilePhoneNumber)
			if len(res) > 0{
				datas["responseNo"] = 0
				datas["password"] = res
			}
		}else{
			datas["responseNo"] = -4
		}
		
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 修改密码
// @Description 修改密码
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	oldPwd			form	string	true	旧密码
// @Param	newPwd			form	string	true	新密码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /pwd/:mobilePhoneNumber [put]
func (u *UserController) ModifyPwd() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	oldPwd := u.Ctx.Request.FormValue("oldPwd")
	newPwd := u.Ctx.Request.FormValue("newPwd")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) && helper.CheckPwdValid(oldPwd) && helper.CheckPwdValid(newPwd) {
		datas["responseNo"] = -1
		if userObj.CheckUserAndPwd(mobilePhoneNumber,oldPwd){
			res2 := userObj.ModifyUserPwd(mobilePhoneNumber,newPwd)
			datas["responseNo"] = res2
		}else{
			datas["responseNo"] = -8
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}

// @Title 验证手机号码是否已注册
// @Description 验证手机号码是否已注册
// @Param	mobilePhoneNumber	path	string	true	手机号码
// @Param	sign			header	string	true	签名
// @Param	pkg			header	string	true	包名
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /exists/:mobilePhoneNumber [get]
func (u *UserController) CheckUserExists() {
	//ini return
	datas := map[string]interface{}{"responseNo": -1}
	//model ini
	var userObj *models.MUser
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Input.Param(":mobilePhoneNumber")
	//check sign
	datas["responseNo"] = u.checkSign(u)
	//检查参数
	if datas["responseNo"] == 0 && helper.CheckMPhoneValid(mobilePhoneNumber) {
		if userObj.CheckUserNameExists(mobilePhoneNumber){
			datas["responseNo"] = -2
		}else{
			datas["responseNo"] = -4
		}
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
	}
	//return
	u.jsonEcho(datas,u)
}