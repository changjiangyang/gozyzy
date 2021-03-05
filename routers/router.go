package routers

import (
	"github.com/astaxie/beego"
	"tsyc/controllers"
	"tsyc/controllers/wechat"
)

func init() {
	beego.Router("/", &controllers.MainController{})
	beego.Router("/wechat", &wechat.WechatController{})
	ns := beego.NewNamespace("/user",
		beego.NSRouter("/getaccesstoken", &controllers.UserController{}, "get:GetAccessToken"),
	)
	beego.AddNamespace(ns)
}
