package provider

import (
	"github.com/google/wire"
	"zhiqu/app/category"
	"zhiqu/app/login"
)

var AppSet = wire.NewSet(NewBeanFactory, category.CategorySet, login.LoginSet)
