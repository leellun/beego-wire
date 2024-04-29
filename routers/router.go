// Package routers 路由配置
package routers

import (
	"github.com/beego/beego/v2/server/web"
	"zhiqu/wire"
)

func init() {
	beanFactory, err := wire.NewApp()
	if err != nil {
		panic(err)
	}
	//这是一种可实现的依赖注入全自动方式
	ns := web.NewNamespace("/v1", func(namespace *web.Namespace) {
		beanFactory.AppFxRouter.Router(namespace)
	})
	web.SetStaticPath("/swagger", "swagger")
	web.AddNamespace(ns)
}
