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
	ui.NewToolBar,
	ui.NewLayout,
), fx.Invoke(setupWindow, showUI))

func newUIApp() (fyne.App, fyne.Preferences) {
	var app = fyneApp.New()
	return app, app.Preferences()
}

func newUIWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Proto")
	window.Resize(fyne.NewSize(600, 800))
	window.CenterOnScreen()
	return window
}

func setupWindow(window fyne.Window, content *fyne.Container) {
	window.SetContent(content)
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
