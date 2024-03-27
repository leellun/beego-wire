package routers

import (
	"github.com/astaxie/beego"
	"zhiqu/controllers"
)

func init() {
	beego.Router("/", &controllers.MainController{})
}
