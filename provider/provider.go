package provider

import (
	"github.com/google/wire"
	"zhiqu/app"
	"zhiqu/infrastructure/database/pgsql"
)

var Set = wire.NewSet(NewBeanFactory, pgsql.NewPgsqlOrm, app.Set)
