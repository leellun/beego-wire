package login

import (
	"github.com/google/wire"
	_ "github.com/google/wire"
)

var LoginSet = wire.NewSet(NewController, NewService, NewDao)
