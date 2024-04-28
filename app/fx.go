package app

import (
	"go.uber.org/fx"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
)

var Module = fx.Module("app", user.Module, category.Module, login.Module)
