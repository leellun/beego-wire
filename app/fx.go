package app

import (
	"go.uber.org/fx"
	"zhiqu/app/user"
)

var Module = fx.Module("app", user.Module)
