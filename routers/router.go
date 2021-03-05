package routers

import (
	"tsyc/controllers"
	"tsyc/controllers/wechat"

	"github.com/astaxie/beego"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wechat", &wechat.WechatController{})
	ns := beego.NewNamespace("/user",
		beego.NSRouter("/getaccesstoken", &controllers.UserController{}, "get:GetAccessToken"),
		beego.NSRouter("/getTemporaryQr", &controllers.UserController{}, "get:GetTemporaryQr"),
	)
	beego.AddNamespace(ns)
}
