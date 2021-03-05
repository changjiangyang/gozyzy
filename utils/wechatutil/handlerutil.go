package wechatutil

import (
	"encoding/xml"
	"tsyc/models/wechatbeans"
	"tsyc/service/wechatService"

	"github.com/astaxie/beego/logs"
)

// ReceiveMessage 微信服务端 消息处理
func ReceiveMessage(msg string) {
	message := wechatbeans.Message{}
	xml.Unmarshal([]byte(msg), &message)
	switch message.MsgType {
	case "text":
		// TODO 文本消息处理
		logs.Info(msg)
		break
	case "image":
		// TODO 图片消息处理
		logs.Info(msg)
		break
	case "voice":
		// TODO 语音消息处理
		break
	case "video":
		// TODO 视频消息处理
		break
	case "shortvideo":
		// TODO 小视频消息处理
		break
	case "location":
		// TODO 地理位置消息处理
		break
	case "link":
		// TODO 链接消息处理
		break
	case "event":
		EventHand(msg)
		break
	}
}

// EventHand 用户事件处理
func EventHand(msg string) {
	message := wechatbeans.Message{}
	xml.Unmarshal([]byte(msg), &message)
	event := message.Event
	switch event {
	case "subscribe":
		// TODO 用户订阅
		logs.Info(message.FromUserName + "关注")
		openid := message.FromUserName
		users := wechatService.GetUserInfo(openid)
		logs.Info(users)
		wechatService.SendTextMessage(openid, "感谢您的关注")
		break
	case "unsubscribe":
		// TODO 取消订阅
		logs.Info(message.FromUserName + "取消关注")
		break
	case "SCAN":
		// TODO 用户已关注时扫描带参数二维码事件
		logs.Info(message.FromUserName + "扫码关注")
		break
	case "LOCATION":
		// TODO 上报地理位置事件
		break
	case "CLICK":
		// TODO 自定义菜单事件
		break
	}
}
