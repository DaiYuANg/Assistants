package cmd

import (
	"context"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"go.uber.org/fx"
	"protocol/assistant/internal/ui"
)

var uiModule = fx.Module("ui", fx.Provide(
	newUIApp,
	newUIWindow,
	ui.NewLayout,
), fx.Invoke(showUI))

func newUIApp() fyne.App {
	var app = fyneApp.New()
	return app
}

func newUIWindow(app fyne.App, content *fyne.Container) fyne.Window {
	window := app.NewWindow("Proto")
	window.SetContent(content)
	window.Resize(fyne.NewSize(600, 800))
	window.CenterOnScreen()
	return window
}

func showUI(lc fx.Lifecycle, window fyne.Window) {
	window.ShowAndRun()
	lc.Append(
		fx.Hook{OnStop: func(ctx context.Context) error {
			window.Close()
			return nil
		}},
	)
}
