package ui

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewContentTab() *container.AppTabs {
	var tab = container.NewAppTabs(
		container.NewTabItem("Tab 1", container.NewVBox(
			widget.NewLabel("Center 1"),
			NewMessageContent(),
		)),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	return tab
}
