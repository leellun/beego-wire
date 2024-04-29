package main

import (
	beego "github.com/beego/beego/v2/server/web"
	"zhiqu/infrastructure/extension/xrecover"
	_ "zhiqu/routers"
)

func main() {
	beego.BConfig.RecoverFunc = xrecover.RecoverPanic
	beego.Run()
}
