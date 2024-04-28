package user

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("user",
	fx.Provide(NewController),
	fx.Provide(NewService),
	fx.Provide(NewDao),
	fx.Invoke(RegisterRouter),
)

func RegisterRouter(router *gin.RouterGroup, controller *Controller) {
	router.GET("/user/:id", controller.GetOne)
}
