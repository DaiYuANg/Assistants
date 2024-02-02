package gui

import (
	"context"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"go.uber.org/fx"
)

type AppResult struct {
	fx.Out
	fyne.App
}

func NewUIApp() AppResult {
	var app = fyneApp.NewWithID("proto.assistant")
	app.SendNotification(fyne.NewNotification("start", "ok"))
	return AppResult{
		App: app,
	}
}

type UIShutdownHookParam struct {
	fx.In
	Lx         fx.Lifecycle
	MainLayout *fyne.Container `name:"mainLayout"`
	App        fyne.App
}

func WindowShutdownHook(param UIShutdownHookParam) {
	param.Lx.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		println("on shutdown")
		param.App.Quit()
		return nil
	}})
}

func ShowUI(lc fx.Lifecycle, app fyne.App, window fyne.Window) {
	window.Show()
	app.Run()
	lc.Append(
		fx.Hook{OnStop: func(ctx context.Context) error {
			window.Close()
			app.Quit()
			return nil
		}},
	)
}
