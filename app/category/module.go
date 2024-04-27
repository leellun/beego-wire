package category

import (
	"github.com/google/wire"
	_ "github.com/google/wire"
)

var CategorySet = wire.NewSet(NewController, NewService, NewDao)
