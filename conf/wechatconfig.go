package conf

import "github.com/astaxie/beego"

// GetAppid 获取appid
func GetAppid() string {
	return beego.AppConfig.String("wechat::appid")
}

// GETTOKEN 获取access_token
var GETTOKEN = "https://api.weixin.qq.com/cgi-bin/token"

// GETUSERINFO 获取用户信息
var GETUSERINFO = "https://api.weixin.qq.com/cgi-bin/user/info"

// SENDMESSAGE 发送客服消息
var SENDMESSAGE = "https://api.weixin.qq.com/cgi-bin/message/custom/send"
