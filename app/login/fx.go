package login

import (
	"github.com/gin-gonic/gin"
	_ "github.com/google/wire"
	"go.uber.org/fx"
)

var Module = fx.Module("login",
	fx.Provide(NewController),
	fx.Provide(NewService),
	fx.Provide(NewDao),
	fx.Invoke(RegisterRouter),
)

func RegisterRouter(router *gin.RouterGroup, controller *Controller) {
	router.GET("/login", controller.Login)
}
