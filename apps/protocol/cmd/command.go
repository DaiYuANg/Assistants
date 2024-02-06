package cmd

import (
	"context"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var app = fx.New(
	fx.Provide(zap.NewExample),
	commandModule,
	infoModule,
	protocolModule,
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
