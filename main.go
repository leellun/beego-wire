package main

import (
	"context"
	"go.uber.org/fx"
	"zhiqu/app"
	"zhiqu/infrastructure/config"
	"zhiqu/infrastructure/database/pgsql"
	"zhiqu/routers"
	_ "zhiqu/routers"
)

func main() {
	config.InitConfig()

	// 定义 FX 应用
	App := fx.New(
		fx.Provide(routers.NewEngine),      // 提供 Gin 引擎
		fx.Provide(routers.NewRouterGroup), // 提供 Gin 引擎
		fx.Provide(pgsql.NewPgsqlDB),       // pgsql
		app.Module,                         // 接口
		fx.Invoke(routers.StartServer),     // 注册路由
	)
	// 启动应用
	if err := App.Start(context.Background()); err != nil {
		panic(err)
	}
	// 阻塞直到程序结束
	<-App.Done()
}
