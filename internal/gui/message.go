package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/fx"
)

type MessageContainerDependency struct {
	fx.In
}

func NewMessageContainer(dependency MessageContainerDependency) *fyne.Container {
	messageContainer := container.NewVBox(
		container.NewAppTabs(
			container.NewTabItem("Tab 1", container.NewVBox(
				widget.NewLabel("Center 1"),
				NewMessageContent(),
			)),
			container.NewTabItem("Tab 2", widget.NewLabel("World!")),
		),
	)
	return messageContainer
}
