package gui

import (
	"fyne.io/fyne/v2"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func NewFxMenu(t interface{}) interface{} {
	return fx.Annotate(
		t,
		fx.ResultTags(`group:"menus"`),
	)
}

func NewFileMenu(logger *zap.Logger) *fyne.Menu {
	return fyne.NewMenu("File",
		fyne.NewMenuItem("Open", func() {
			logger.Info("test", zap.String("open", "open"))
			println("open")
		}),
	)
}

func NewHelpMenu() *fyne.Menu {
	return fyne.NewMenu("Help",
		fyne.NewMenuItem("About", func() {
			println("about")
		}),
	)
}

func NewMainMenu(menus []*fyne.Menu) *fyne.MainMenu {
	return fyne.NewMainMenu(menus...)
}
