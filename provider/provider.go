package provider

import (
	"github.com/google/wire"
	"zhiqu/app/category"
	"zhiqu/app/login"
	"zhiqu/app/user"
	"zhiqu/infrastructure/database/pgsql"
)

var AppSet = wire.NewSet(NewBeanFactory, pgsql.NewPgsqlOrm, category.CategorySet, login.Set, user.Set)
