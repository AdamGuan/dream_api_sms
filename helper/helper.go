package helper

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"os"
	"strconv"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"crypto/md5"
	"net/url"
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
func CurlLeanCloud(requestUri string, method string, requestData map[string]string) map[string]interface{} {
	geturl := requestUri
	req, _ := http.NewRequest(method, geturl, nil)
	data, _ := json.Marshal(requestData)
	req.Body = ioutil.NopCloser(strings.NewReader(string(data)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("User-Agent", "SSTS Browser/1.0")
	req.Header.Add("X-AVOSCloud-Application-Id", "l7dn2hfzry3yxdtfhnhciy0jt00dqd3ysbees8pjdhonwjb4")
	req.Header.Add("X-AVOSCloud-Application-Key", "76nib8wzexxcgaccrl9e4955ll8q11w1jwq36ofpx6u340q2")
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
func CheckSign(requestParames url.Values, key string) bool {
	//key
	//key := "test"
	//计算sign
	if signarray, ok := requestParames["sign"]; ok {
		if sign := signarray[0]; len(signarray) > 0 {
			requestParames.Del("sign")
			str := requestParames.Encode() + "&key=" + key
			sign2 := fmt.Sprintf("%x\r\n", md5.Sum([]byte(str)))
			sign2 = strings.TrimSpace(sign2)
			sign = strings.TrimSpace(sign)

			if string(sign) == string(sign2) {
				return true
			}
		}
	}

	return false
}
