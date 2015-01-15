package models

import (
	"dream_api_sms/helper"
)

func init() {
}

type MSms struct {
}

//get a msm
func (u *MSms) GetMsm(mobilePhoneNumber string,appId string,appKey string,appName string,appTemplate string) map[string]interface{} {
	return helper.CurlLeanCloud("https://leancloud.cn/1.1/requestSmsCode","POST",map[string]string{"mobilePhoneNumber": mobilePhoneNumber,"template":appTemplate,"appname":appName},appId,appKey);
}

//valid a msm
func (u *MSms) ValidMsm(num string,mobilePhoneNumber string,appId string,appKey string) map[string]interface{} {
	return helper.CurlLeanCloud("https://leancloud.cn/1.1/verifySmsCode/"+num+"?mobilePhoneNumber="+mobilePhoneNumber,"POST",map[string]string{},appId,appKey);
}