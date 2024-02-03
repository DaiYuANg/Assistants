package cmd

import (
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

var commandModule = fx.Module("command", fx.Provide(NewCommand))

func NewCommand() cobra.Command {
	var cmd = cobra.Command{Use: "Protocol",
		SilenceErrors: true,
		SilenceUsage:  true,
		RunE: func(cmd *cobra.Command, args []string) error {
			//return app.Start(context.Background())
			return nil
		},
	}
	return cmd
}
