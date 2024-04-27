package login

import (
	"github.com/google/wire"
	_ "github.com/google/wire"
)

var Set = wire.NewSet(NewController, NewService, NewDao)
