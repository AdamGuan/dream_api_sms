package helper

import (
	"crypto/md5"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	//"net/url"
	"time"
	"regexp"
)

var MyLog *logs.BeeLogger

func init() {
	//初始化log
	if beego.RunMode == "dev" {
		MyLog = logs.NewLogger(10000)
		MyLog.SetLogger("file", `{"filename":"log.log"}`)
		MyLog.EnableFuncCallDepth(true)
		//MyLog.Debug("debug test1")
		// MyLog.Error("error")
	}
}

//类型转化 string  to int
func StrToInt(str string) int {
	i, _ := strconv.Atoi(str)
	return i
}

//类型转化 string  to float64
func StrToFloat64(str string) float64 {
	f, _ := strconv.ParseFloat(str, 64)
	return f
}

//类型转化 int to string
func IntToString(i int) string {
	return fmt.Sprintf("%d", i)
}

//类型转化 int64 to string
func Int64ToString(i int64) string {
	return fmt.Sprintf("%d", i)
}

//获取16位Guid
func GetGuid() string {
	f, _ := os.OpenFile("/dev/urandom", os.O_RDONLY, 0)
	b := make([]byte, 16)
	f.Read(b)
	f.Close()
	uuid := fmt.Sprintf("%x", b)
	return uuid[0:16]
}

//创建密码
func CreatePwd(num int) string {
	return GetGuid()[0:num]
}

//leanCloud curl
func CurlLeanCloud(requestUri string, method string, requestData map[string]string, appId string, appKey string) map[string]interface{} {
	geturl := requestUri
	req, _ := http.NewRequest(method, geturl, nil)
	data, _ := json.Marshal(requestData)
	req.Body = ioutil.NopCloser(strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "SSTS Browser/1.0")
	//req.Header.Add("X-AVOSCloud-Application-Id", "l7dn2hfzry3yxdtfhnhciy0jt00dqd3ysbees8pjdhonwjb4")
	req.Header.Add("X-AVOSCloud-Application-Id", appId)
	//req.Header.Add("X-AVOSCloud-Application-Key", "76nib8wzexxcgaccrl9e4955ll8q11w1jwq36ofpx6u340q2")
	req.Header.Add("X-AVOSCloud-Application-Key", appKey)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		},
	}

	resp, _ := client.Do(req)
	defer resp.Body.Close()
	bodyByte, _ := ioutil.ReadAll(resp.Body)
	p := map[string]interface{}{}
	json.Unmarshal(bodyByte, &p)
	return p
}

//检查签名
func CheckSign(sign string, pkg string) bool {
	//sign = timestamp+md5(pkgname+timestamp)
	if len(sign) == 46 && len(pkg) > 0 {
		timestamp := sign[0:14]
		//检测是否超时
		if beego.RunMode != "dev"{
			nowTime, _ := strconv.Atoi(time.Now().Format("20060102150405"))
			requestTime, _ := strconv.Atoi(timestamp)
			timedistince := nowTime - requestTime
			fmt.Println(timedistince)
			if timedistince > 60*5 {
				return false
			}
		}

		sign = sign[14:]
		str := pkg + timestamp
		sign2 := fmt.Sprintf("%x\r\n", md5.Sum([]byte(str)))
		sign2 = strings.TrimSpace(sign2)
		if string(sign) == string(sign2) {
			return true
		}
	}
	return false
}

//手机号码有效性验证
func CheckMPhoneValid(phone string)bool{
	matched, err := regexp.MatchString("^1[3|4|5|8][0-9]{9}$", phone)
	if err == nil && matched{
		return true
	}
	return false
}

//密码有效性验证
func CheckPwdValid(pwd string)bool{
	//pwd = strings.TrimSpace(pwd)
	matched, err := regexp.MatchString("^\\w{6,40}$", pwd)
	if err == nil && matched{
		return true
	}
	return false
}
