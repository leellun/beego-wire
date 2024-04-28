package provider

import (
	"github.com/google/wire"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
)

var AppSet = wire.NewSet(NewBeanFactory, category.CategorySet, login.Set, user.Set)
