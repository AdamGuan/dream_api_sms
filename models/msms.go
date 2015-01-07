package models

import (
	//"crypto/md5"
	//"fmt"
	//"github.com/astaxie/beego/orm"
	//_ "github.com/go-sql-driver/mysql"
	"dream_api_sms/helper"
)

func init() {
}

type MSms struct {
}

//get a msm
func (u *MSms) GetMsm(mobilePhoneNumber string) map[string]interface{} {
	return helper.CurlLeanCloud("https://leancloud.cn/1.1/requestSmsCode","POST",map[string]string{"mobilePhoneNumber": mobilePhoneNumber});
}

//valid a msm
func (u *MSms) ValidMsm(num string,mobilePhoneNumber string) map[string]interface{} {
	return helper.CurlLeanCloud("https://leancloud.cn/1.1/verifySmsCode/"+num+"?mobilePhoneNumber="+mobilePhoneNumber,"POST",map[string]string{});
}