package main

import (
	"github.com/astaxie/beego"
	_ "tsyc/routers"
	"tsyc/task"
)

func main() {
	task.FlashToken()
	beego.Run()
}
