package cmd

import (
	"context"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

var app = fx.New(
	uiModule,
	errorModule,
	fx.Provide(zap.NewExample),
)

var RootCmd = cobra.Command{Use: "ProtocolAssistant",
	RunE: func(cmd *cobra.Command, args []string) error {
		return app.Start(context.Background())
	}}

func Execute() error {
	return RootCmd.Execute()
}
