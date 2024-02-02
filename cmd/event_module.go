package cmd

import (
	"go.uber.org/fx"
	"protocol/assistant/internal/event"
)

var eventModule = fx.Module("eventModule",
	fx.Provide(
		fx.Annotate(event.NewGlobalEventEmitter, fx.ResultTags(`name:"globalEventMitter"`)),
		fx.Annotate(event.NewGUIEventEmitter, fx.ResultTags(`name:"guiEventEmitter"`)),
		fx.Annotate(event.NewBackendEventEmitter, fx.ResultTags(`name:"BackendEventEmitter"`)),
	),
)
