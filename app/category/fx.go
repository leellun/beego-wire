package category

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
	namespace.Get("/category", r.controller.GetAll)
	namespace.Get("/category/:id", r.controller.GetOne)
	namespace.Post("/category", r.controller.Post)
	namespace.Put("/category/:id", r.controller.Put)
	namespace.Delete("/category/:id", r.controller.Delete)
}

var Set = wire.NewSet(NewController, NewService, NewDao, NewAppRouter)
