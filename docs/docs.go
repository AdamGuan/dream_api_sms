package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

const (
    Rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/sms","description":"短信\n"},{"path":"/user","description":"用户\n"}],"info":{"title":"用户系统 API"}}`
    Subapi string = `{"/sms":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/sms","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/smsvalid/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"POST","nickname":"短信验证码验证","type":"","summary":"短信验证码验证","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"\"手机号码\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"num","description":"\"验证码\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"\"签名\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"\"包名\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/register/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"发送一条短信验证码(注册时)","type":"","summary":"发送一条短信验证码(注册时)","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/resetpwd/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"发送一条短信验证码(重置密码时)","type":"","summary":"发送一条短信验证码(重置密码时)","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/pwd/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"发送一条短信验证码(找回密码时)","type":"","summary":"发送一条短信验证码(找回密码时)","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]}],"models":{"MResp":{"id":"MResp","properties":{"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}}}},"/user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/register","description":"","operations":[{"httpMethod":"POST","nickname":"注册","type":"","summary":"注册","parameters":[{"paramType":"form","name":"mobilePhoneNumber","description":"\"手机号码\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"pwd","description":"\"密码\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"\"签名\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"\"包名\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/resetpwd","description":"","operations":[{"httpMethod":"PUT","nickname":"重置密码","type":"","summary":"重置密码","parameters":[{"paramType":"form","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"pwd","description":"密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/login/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"登录验证","type":"","summary":"登录验证","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"pwd","description":"密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/pwd/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"GET","nickname":"找回密码","type":"","summary":"找回密码","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"\"手机号码\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"\"签名\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"\"包名\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MFindPwdResp","responseModel":"MFindPwdResp"},{"code":401,"message":"无权访问","responseModel":""}]}]},{"path":"/pwd/:mobilePhoneNumber","description":"","operations":[{"httpMethod":"PUT","nickname":"修改密码","type":"","summary":"修改密码","parameters":[{"paramType":"path","name":"mobilePhoneNumber","description":"手机号码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"oldPwd","description":"旧密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"form","name":"newPwd","description":"新密码","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"sign","description":"签名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"header","name":"pkg","description":"包名","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.MResp","responseModel":"MResp"},{"code":401,"message":"无权访问","responseModel":""}]}]}],"models":{"MFindPwdResp":{"id":"MFindPwdResp","properties":{"password":{"type":"string","description":"","format":""},"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}},"MResp":{"id":"MResp","properties":{"responseMsg":{"type":"string","description":"","format":""},"responseNo":{"type":"int","description":"","format":""}}}}}}`
    BasePath string= "/v1"
)

var rootapi swagger.ResourceListing
var apilist map[string]*swagger.ApiDeclaration

func init() {
	err := json.Unmarshal([]byte(Rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(Subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = BasePath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
