package content

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"protocol/assistant/internal/gui"
)

func NewContentTab() *container.AppTabs {
	var tab = container.NewAppTabs(
		container.NewTabItemWithIcon("Tab 1", theme.CancelIcon(), container.NewVBox(
			widget.NewLabel("Center 1"),
			gui.NewMessageContent(),
		)),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
	)

	return tab
}
