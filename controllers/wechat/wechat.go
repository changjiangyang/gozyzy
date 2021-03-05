package wechat

import (
	"tsyc/utils/wechatutil"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

type WechatController struct {
	beego.Controller
}

func (c *WechatController) Get() {
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	echostr := c.GetString("echostr")
	tok := beego.AppConfig.String("wechat::token")
	b := wechatutil.AuthMessage(tok, timestamp, nonce, signature)
	if b {
		c.Ctx.WriteString(echostr)
	} else {
		c.Ctx.WriteString("false")
	}

}

// Post 微信消息处理
func (c *WechatController) Post() {
	sff := string(c.Ctx.Input.RequestBody)
	signature := c.GetString("signature")
	timestamp := c.GetString("timestamp")
	nonce := c.GetString("nonce")
	tok := beego.AppConfig.String("wechat::token")
	b := wechatutil.AuthMessage(tok, timestamp, nonce, signature)
	if b {
		logs.Info(sff + "%%%%sss")
		wechatutil.ReceiveMessage(sff)
		c.Ctx.WriteString("")
	} else {
		c.Ctx.WriteString("签名验证错误")
	}

}
