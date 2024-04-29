package login

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/google/wire"
)

type FxRouter struct {
	controller *Controller
}

func NewAppRouter(controller *Controller) *FxRouter {
	return &FxRouter{controller}
}
func (r *FxRouter) Router(namespace *web.Namespace) {
	namespace.Post("/login", r.controller.Login)
}

var Set = wire.NewSet(NewController, NewService, NewDao, NewAppRouter)
