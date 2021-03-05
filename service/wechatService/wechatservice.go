package wechatService

import (
	"encoding/json"
	"fmt"
	"tsyc/conf"
	"tsyc/models/wechatbeans"
	"tsyc/utils/cacheutil"
	"tsyc/utils/httputil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

// GetAccessToken 获取access_token
func GetAccessToken() string {
	appid := beego.AppConfig.String("wechat::appid")
	return cacheutil.GetToken(appid)
}

// GetUserInfo 从微信服务端获取 用户信息
func GetUserInfo(openid string) wechatbeans.Users {
	var url = conf.GETUSERINFO
	token := GetAccessToken()
	var param map[string]string
	param = make(map[string]string)
	param["access_token"] = token
	param["openid"] = openid
	param["lang"] = "zh_CN"
	user := new(wechatbeans.Users)
	res, err := httputil.Get(url, param)
	if err != nil {
		logs.Error(err)
	} else {
		fmt.Println(res)
		//var usermap map[string]interface{}
		//usermap := make(map[string]interface{})

		err := json.Unmarshal([]byte(res), &user)
		if err == nil {
			if user.Subscribe == 0 {
				// 用户没有关注公众号 不做处理
			} else {
				// TODO 关注公众号处理
				fmt.Println(user)
			}
		} else {
			logs.Error(err)
		}

	}
	return *user
}

// SendTextMessage 发送文本消息
func SendTextMessage(openid string, context string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "text"
	content := make(map[string]string)
	content["content"] = context
	param["text"] = content
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}
