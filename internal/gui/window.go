package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"go.uber.org/fx"
)

func NewUIWindow(app fyne.App) fyne.Window {
	window := app.NewWindow("Proto")
	window.Resize(fyne.NewSize(800, 800))
	window.CenterOnScreen()
	window.SetCloseIntercept(func() {
		fmt.Println("close")
		quitConfirm := dialog.NewConfirm("Quit", "Realy", func(b bool) {
			if b {
				window.Close()
			}
		}, window)
		quitConfirm.Show()
	})
	return window
}

type SetupWindowParam struct {
	fx.In
	Window   fyne.Window
	Content  *fyne.Container `name:"mainLayout"`
	MainMenu *fyne.MainMenu
}

func SetupWindow(param SetupWindowParam) {
	param.Window.SetContent(param.Content)
	param.Window.SetMainMenu(param.MainMenu)
}

func NewUIChan() chan struct{} {
	return make(chan struct{})
}
