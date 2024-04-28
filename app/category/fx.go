package category

import "go.uber.org/fx"

var Module = fx.Module("category",
	fx.Provide(NewController),
	fx.Provide(NewService),
	fx.Provide(NewController),
)
