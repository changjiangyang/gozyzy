package controllers

import (
	"tsyc/service/wechatService"
	"tsyc/utils/cacheutil"

	"github.com/astaxie/beego"
)

// UserController 用户接口
type UserController struct {
	beego.Controller
}

// GetAccessToken 获取AccessToken
func (user *UserController) GetAccessToken() {
	appid := beego.AppConfig.String("wechat::appid")
	accessToken := cacheutil.GetToken(appid)
	user.Ctx.WriteString(accessToken)
}

// GetTemporaryQr 获取临时二维码
func (c *UserController) GetTemporaryQr() {
	res := wechatService.GetTemporaryQr(600, "o1kII5zxIJCcv8E_ry_GODJi-jqc")
	c.Data["json"] = res
	c.ServeJSON()
}
