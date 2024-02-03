package cmd

import (
	"context"
	"fx_module"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var app = fx.New(
	fx.Provide(zap.NewExample),
	commandModule,
	infoModule,
	fx_module.GlobalEventModule,
	eventModule,
	configModule,
	schedulerModule,
	databaseModule,
	contextModule,
	errorModule,
	uiModule,
)

func Execute() error {
	return app.Start(context.Background())
}
