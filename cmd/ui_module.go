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
	fx.Annotate(ui.NewToolBar, fx.ResultTags(`name:"toolBar"`)),
	fx.Annotate(ui.NewLayout, fx.ResultTags(`name:"layout"`)),
	fx.Annotate(ui.NewStatusBar, fx.ResultTags(`name:"statusBar"`)),
), fx.Invoke(setupWindow, showUI, windowShutdownHook))

type UIAppResult struct {
	fx.Out
	fyne.App
	fyne.Preferences
}

func newUIApp() UIAppResult {
	var app = fyneApp.NewWithID("proto.assistant")
	app.SendNotification(fyne.NewNotification("start", "ok"))
	return UIAppResult{
		App:         app,
		Preferences: app.Preferences(),
	}
}

func newUIWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Proto")
	window.Resize(fyne.NewSize(600, 800))
	window.CenterOnScreen()
	return window
}

type SetupWindowParam struct {
	fx.In
	Window  fyne.Window
	Content *fyne.Container `name:"layout"`
}

func setupWindow(param SetupWindowParam) {
	param.Window.SetContent(param.Content)
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

type WindowShutdownHookParam struct {
	fx.In
	Lx          fx.Lifecycle
	MainLayout  *fyne.Container `name:"layout"`
	Preferences fyne.Preferences
}

func windowShutdownHook(param WindowShutdownHookParam) {
	param.Lx.Append(fx.Hook{OnStop: func(ctx context.Context) error {
		println("on shutdown")
		return nil
	}})
}
