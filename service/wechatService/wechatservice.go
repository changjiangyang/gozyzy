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

// SendImgMessage 发送图片消息
func SendImgMessage(openid string, mediaid string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "image"
	content := make(map[string]string)
	content["media_id"] = mediaid
	param["image"] = content
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}

// SendVoiceMessage 发送语音消息
func SendVoiceMessage(openid string, mediaid string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "voice"
	content := make(map[string]string)
	content["media_id"] = mediaid
	param["image"] = content
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}

// SendVideoMessage 发送视频消息
func SendVideoMessage(openid string, media_id string, thumb_media_id string, title string, desc string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "video"
	video := make(map[string]string)
	video["media_id"] = media_id
	video["thumb_media_id"] = thumb_media_id
	video["title"] = title
	video["description"] = desc
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}

// SendMusicMessage 发送音乐消息
func SendMusicMessage(openid string, title string, musicurl string, hqurl string, media_id string, desc string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "music"
	music := make(map[string]string)
	music["title"] = title
	music["description"] = desc
	music["musicurl"] = musicurl
	music["hqmusicurl"] = hqurl
	music["thumb_media_id"] = media_id
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}

// SendNewsMessage 发送图文消息（点击跳转到外链） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
func SendNewsMessage(openid string, title string, desc string, newsurl string, picurl string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "news"
	news := make(map[string]string)
	news["title"] = title
	news["description"] = desc
	news["url"] = newsurl
	news["picurl"] = picurl
	articles := [1]interface{}{news}
	param["news"] = articles
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}

// SendMpnewsMessage 发送图文消息（点击跳转到图文消息页面） 图文消息条数限制在1条以内，注意，如果图文数超过1，则将会返回错误码45008。
func SendMpnewsMessage(openid string, mediaid string) {
	url := conf.SENDMESSAGE
	url = url + "?access_token=" + GetAccessToken()
	param := make(map[string]interface{})
	param["touser"] = openid
	param["msgtype"] = "mpnews"
	mpnews := make(map[string]string)
	mpnews["media_id"] = mediaid
	res, err := httputil.SendPostJson(url, param, nil)
	if err == nil {
		fmt.Println(res)
	}
}
