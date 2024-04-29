package app

import (
	"github.com/beego/beego/v2/server/web"
	"github.com/google/wire"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
)

type FxRouter struct {
	categoryFxRouter *category.FxRouter
	loginFxRouter    *login.FxRouter
	userFxRouter     *user.FxRouter
}

func NewAppRouter(categoryFxRouter *category.FxRouter,
	loginFxRouter *login.FxRouter,
	userFxRouter *user.FxRouter) *FxRouter {
	return &FxRouter{categoryFxRouter, loginFxRouter, userFxRouter}
}
func (r *FxRouter) Router(namespace *web.Namespace) {
	r.categoryFxRouter.Router(namespace)
	r.loginFxRouter.Router(namespace)
	r.userFxRouter.Router(namespace)
}

var Set = wire.NewSet(NewAppRouter, category.Set, login.Set, user.Set)
