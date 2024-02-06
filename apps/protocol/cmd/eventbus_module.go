package cmd

import (
	"eventbus"
	"go.uber.org/fx"
)

var eventBusModule = fx.Module("eventBus", fx.Provide(eventbus.NewEventBus[any]()))
