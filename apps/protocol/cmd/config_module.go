package cmd

import (
	"asssistant/protocol/internal/config"
	"go.uber.org/fx"
)

var configModule = fx.Module("configModule",
	fx.Provide(config.NewConfig),
	fx.Invoke(config.ListenConfig),
)
