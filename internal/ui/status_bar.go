package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func NewStatusBar() *fyne.Container {
	var statusBar = container.NewHBox(
		widget.NewLabel("Left 1"),
		widget.NewLabel("Left 2"),
		widget.NewProgressBarInfinite(),
	)
	return statusBar
}
