package cmd

import (
	"github.com/go-co-op/gocron/v2"
	"go.uber.org/fx"
)

var schedulerModule = fx.Module("schedulerModule", fx.Provide(newScheduler))

func newScheduler() *gocron.Scheduler {
	s, err := gocron.NewScheduler()
	if err != nil {
		panic(err)
	}
	return &s
}
