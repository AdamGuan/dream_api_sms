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
		res2 := userObj.AddUser(mobilePhoneNumber,pwd)
		datas["responseNo"] = res2
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
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
		res2 := userObj.ModifyUserPwd(mobilePhoneNumber,pwd)
		datas["responseNo"] = res2
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -1
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
	}else if datas["responseNo"] == 0{
		datas["responseNo"] = -5
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
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
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
	if datas["responseNo"] == 0 && len(mobilePhoneNumber) > 0 && len(oldPwd) > 0 && len(newPwd) > 0 {
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
	if datas["responseNo"] == -6 || datas["responseNo"] == -7 {
		u.Ctx.ResponseWriter.Header().Set("Content-Type", "application/json; charset=utf-8")
		u.Ctx.ResponseWriter.WriteHeader(http.StatusUnauthorized)
	} 
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}