// Package routers @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"zhiqu/infrastructure/config"
)

func NewAPIRouterGroup() *gin.RouterGroup {
	engine := gin.Default()
	routerGroup := engine.Group("/api") // 创建一个路由组
	return routerGroup
}
func StartServer(apiRouterGroup *gin.RouterGroup) {
	// 这里可以对路由组进行进一步操作，例如添加中间件等
	// 打印启动信息
	fmt.Println("Server is starting...")

	// 获取监听地址和端口
	listenAddr := config.AppConfig.App.Host
	port := config.AppConfig.App.Port

	// 启动服务器并监听指定地址和端口
	addr := fmt.Sprintf("%s:%d", listenAddr, port)
	go func() {
		if err := gin.Default().Run(addr); err != nil {
			fmt.Printf("Server failed to start: %v\n", err)
		}
	}()

	// 打印监听信息
	fmt.Printf("Server is now listening on %s\n", addr)
}
