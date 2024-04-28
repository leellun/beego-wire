package category

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

var Module = fx.Module("category",
	fx.Provide(NewController),
	fx.Provide(NewService),
	fx.Provide(NewDao),
	fx.Invoke(RegisterRouter),
)

func RegisterRouter(router *gin.RouterGroup, controller *Controller) {
	router.GET("/category/:id", controller.GetOne)
}
