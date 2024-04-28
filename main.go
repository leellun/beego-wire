package main

import (
	"zhiqu/infrastructure/config"
	_ "zhiqu/routers"
)

func main() {
	config.InitConfig()
	//_, err := wire.NewApp()
	//if err != nil {
	//	panic(err)
	//}
	//beego.BConfig.RecoverFunc = xrecover.RecoverPanic
	//beego.ErrorController(controller.ErrorController{})
	//beego.Run()
}
