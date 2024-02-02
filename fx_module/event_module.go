package fx_module

import (
	"github.com/kataras/go-events"
	"go.uber.org/fx"
)

var GlobalEventModule = fx.Module("GlobalEventModule", fx.Provide(NewGlobalEventEmitter))

func NewGlobalEventEmitter() events.EventEmmiter {
	e := events.New()
	return e
}
