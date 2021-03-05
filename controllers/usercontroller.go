package controllers

import (
	"github.com/astaxie/beego"
	"tsyc/utils/cacheutil"
)

type UserController struct {
	beego.Controller
}

func (user *UserController) GetAccessToken() {
	appid := beego.AppConfig.String("wechat::appid")
	accessToken := cacheutil.GetToken(appid)
	user.Ctx.WriteString(accessToken)
}
