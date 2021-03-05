package task

import (
	"encoding/json"
	"fmt"
	"tsyc/conf"
	"tsyc/utils/cacheutil"
	_ "tsyc/utils/cacheutil"
	"tsyc/utils/httputil"

	"github.com/astaxie/beego"
	_ "github.com/astaxie/beego/cache"
	"github.com/astaxie/beego/toolbox"
)

// FlashToken 定时刷新token
func FlashToken() {

	tk1 := toolbox.NewTask("tk1",
		"0 0 0/1 * * *",
		func() error {
			fmt.Println("tk1")
			appid := beego.AppConfig.String("wechat::appid")
			server := beego.AppConfig.String("wechat::server")
			par := map[string]string{"appid": appid, "secret": server, "grant_type": "client_credential"}
			res, err := httputil.Get(conf.GETTOKEN, par)
			if err == nil {
				fmt.Println(res)
				var mapResult map[string]interface{}
				err := json.Unmarshal([]byte(res), &mapResult)
				fmt.Println(mapResult)
				if err == nil {
					cacheutil.PutToken(appid, mapResult["access_token"])
				} else {
					fmt.Println(err)
				}
			}
			return nil
		})

	toolbox.AddTask("tk1", tk1)
	toolbox.StartTask()

}
