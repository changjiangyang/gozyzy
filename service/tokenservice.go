package service

import (
	"fmt"
	"github.com/astaxie/beego/cache"
)

func GetAssessToken(appid string) string {
	bm, err := cache.NewCache("redis", `{"key":"default","conn":"6379", "password":"", "dbNum":"0"}`)
	fmt.Println(err)
	if err == nil {
		access_token := fmt.Sprintf("%v", bm.Get(appid))
		return access_token
	} else {
		return ""
	}
}
