package cmd

import (
	"go.uber.org/fx"
	"protocol/internal/config"
)

var configModule = fx.Module("configModule",
	fx.Provide(config.NewConfig),
	fx.Invoke(config.ListenConfig),
)
