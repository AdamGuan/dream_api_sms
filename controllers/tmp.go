package controllers

import (
	"dream_api_sms/models"
	"github.com/astaxie/beego"
	"dream_api_sms/helper"
)

//用户
type TmpController struct {
	beego.Controller
}

// @Title 清空全部用户数据(临时用)
// @Description 清空全部用户数据(临时用)
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /alluser [delete]
func (u *TmpController) DeleteAllUser() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//model ini
	var tmpObj *models.MTmp
	tmpObj.DeleteAllUser()
	//return
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}

// @Title 清空指定用户数据(临时用)
// @Description 清空指定用户数据(临时用)
// @Param	mobilePhoneNumber			query	string	true	手机号码
// @Success	200 {object} models.MResp
// @Failure 401 无权访问
// @router /user [delete]
func (u *TmpController) DeleteUser() {
	//ini return
	datas := map[string]interface{}{"responseNo": 0}
	//parse request parames
	u.Ctx.Request.ParseForm()
	mobilePhoneNumber := u.Ctx.Request.FormValue("mobilePhoneNumber")
	//model ini
	var tmpObj *models.MTmp
	tmpObj.DeleteUser(mobilePhoneNumber)
	//return
	datas["responseMsg"] = models.ConfigMyResponse[helper.IntToString(datas["responseNo"].(int))]
	u.Data["json"] = datas
	u.ServeJson()
}