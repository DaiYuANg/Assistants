package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/fx"
)

type LayoutDependency struct {
	fx.In
	Window      fyne.Window
	Toolbar     *widget.Toolbar  `name:"toolBar"`
	StatusBar   *fyne.Container  `name:"statusBar"`
	MainContent *container.Split `name:"mainContent"`
}

func NewMainLayout(dependency LayoutDependency) *fyne.Container {
	mainLayout :=
		container.NewBorder(
			// top
			dependency.Toolbar,
			// bottom
			dependency.StatusBar,
			// left
			nil,
			// right
			nil,
			dependency.MainContent,
		)
	return mainLayout
}

func NewMessageContent() *fyne.Container {
	text := widget.NewRichTextWithText("test")
	text.Resize(fyne.NewSize(50, 100))
	scroll := container.NewScroll(
		text,
	)
	scroll.Resize(fyne.NewSize(50, 600))
	scroll.Resize(fyne.NewSize(500, 500))
	return container.NewVBox(
		scroll,
	)
}
