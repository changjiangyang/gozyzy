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

// SENDTEMPLATE 发送模板消息
var SENDTEMPLATE = "https://api.weixin.qq.com/cgi-bin/message/template/send"

// GETQRCODE 获取二维码
var GETQRCODE = "https://api.weixin.qq.com/cgi-bin/qrcode/create"
