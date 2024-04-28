// Package routers 路由信息
package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"runtime/debug"
	"zhiqu/infrastructure/config"
	"zhiqu/infrastructure/constant"
	"zhiqu/infrastructure/extension/errox"
	"zhiqu/infrastructure/extension/logz"
	"zhiqu/infrastructure/response"
)

func NewEngine() *gin.Engine {
	engine := gin.Default()
	engine.Use(Middleware())
	return engine
}
func NewRouterGroup(engine *gin.Engine) *gin.RouterGroup {
	routerGroup := engine.Group("/api") // 创建一个路由组
	return routerGroup
}

func StartServer(engine *gin.Engine, routerGroup *gin.RouterGroup) {
	fmt.Println("Server is starting...")
	// 获取监听地址和端口
	listenAddr := config.AppConfig.App.Host
	port := config.AppConfig.App.Port

	// 启动服务器并监听指定地址和端口
	addr := fmt.Sprintf("%s:%d", listenAddr, port)
	go func() {
		if err := engine.Run(addr); err != nil {
			fmt.Printf("Server failed to start: %v\n", err)
		}
	}()

	// 打印监听信息
	fmt.Printf("Server is now listening on %s\n", addr)
}

// Middleware 是一个 Gin 中间件函数
func Middleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			err := recover()
			if err == nil {
				return
			}
			switch v := err.(type) {
			case errox.Exception:
				// 符合预期的错误，可以直接返回给客户端
				logz.Warn(fmt.Sprintf("业务错误：%s", v.Message))
				debug.PrintStack()
				if v.Code != 0 {
					ctx.Status(v.Code)
				}
				ctx.JSON(v.Code, response.Error(v.Message))
			case error:
				// 一律返回服务器错误，避免返回堆栈错误给客户端，实际还可以针对系统错误做其他处理
				logz.Error(fmt.Sprintf("意外错误：%s", v.Error()))
				debug.PrintStack()
				ctx.Status(int(constant.INTERNAL_SERVER_ERROR))
				ctx.JSON(int(constant.INTERNAL_SERVER_ERROR), response.Error(err.(string)))
				break
			default:
				// 一律返回服务器错误，避免返回堆栈错误给客户端，实际还可以针对系统错误做其他处理
				logz.Error(fmt.Sprintf("意外错误：%v", v))
				debug.PrintStack()
				ctx.Status(int(constant.INTERNAL_SERVER_ERROR))
				ctx.JSON(int(constant.INTERNAL_SERVER_ERROR), response.Error(err.(string)))
				break
			}
		}()
		ctx.Next()
	}
}
