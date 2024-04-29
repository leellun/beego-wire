package user

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
	namespace.Get("/user", r.controller.GetAll)
	namespace.Get("/user/:id", r.controller.GetOne)
	namespace.Post("/user", r.controller.Post)
	namespace.Put("/user/:id", r.controller.Put)
	namespace.Delete("/user/:id", r.controller.Delete)
}

var Set = wire.NewSet(NewController, NewService, NewDao, NewAppRouter)
