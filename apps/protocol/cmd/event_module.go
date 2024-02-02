package cmd

import (
	"asssistant/protocol/internal/event"
	"go.uber.org/fx"
)

var eventModule = fx.Module("eventModule",
	fx.Provide(
		fx.Annotate(event.NewGUIEventEmitter, fx.ResultTags(`name:"guiEventEmitter"`)),
		fx.Annotate(event.NewBackendEventEmitter, fx.ResultTags(`name:"BackendEventEmitter"`)),
	),
)
