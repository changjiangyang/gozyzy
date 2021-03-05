package cacheutil

import (
	"fmt"
	"github.com/astaxie/beego/cache"
	_ "github.com/astaxie/beego/cache/redis"
	"github.com/astaxie/beego/logs"
	"time"
)

var err error = nil
var bc cache.Cache = nil

func init() {
	if bc == nil {
		bc, err = cache.NewCache("redis", `{"key":"default","conn":"127.0.0.1:6379", "password":"", "dbNum":"0"}`)
		if err != nil {
			logs.Error(err)
		}
	}

}

func PutToken(key string, value interface{}) {
	fmt.Println("----------------------------------")
	fmt.Println(key)
	fmt.Println(value)
	bc.Put(key, value, 7200*time.Second)
	accresstoken := bc.Get(key)
	fmt.Println(accresstoken)
	fmt.Println("----------------------------------")
}

func GetToken(key string) string {
	accresstoken := bc.Get(key)
	return string(accresstoken.([]byte))
}
