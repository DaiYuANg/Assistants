package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func NewStatusBar() *fyne.Container {
	var statusBar = container.NewHBox(
		widget.NewIcon(theme.AccountIcon()),
		widget.NewProgressBarInfinite(),
	)
	return statusBar
}
