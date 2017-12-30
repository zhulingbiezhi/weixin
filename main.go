package main

import (
	_ "weixin/models"
	_ "weixin/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

func main() {
	beego.BeeLogger.SetLevel(logs.LevelError)
	beego.Run()
}
