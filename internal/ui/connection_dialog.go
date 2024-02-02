package ui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	"protocol/assistant/internal/component"
)

func NewCreateConnectionDialog(window fyne.Window) {
	var button = widget.NewButton("test", func() {
		println("test")
	})
	button.Disable()
	var label = widget.NewLabelWithStyle("Protocol", fyne.TextAlignLeading, fyne.TextStyle{
		Bold:      true,
		Italic:    false,
		Monospace: true,
		Symbol:    true,
		TabWidth:  0,
	})
	var sp = widget.NewSeparator()
	sp.ExtendBaseWidget(label)
	dl := dialog.NewCustomConfirm(
		"Create connection",
		"Save",
		"cancel",
		container.NewVBox(
			label,
			sp,
			widget.NewSelect([]string{"TCP", "UDP"}, func(s string) {

			}),
			component.NewButton("test", func(button *widget.Button) {
				button.SetText("test change button")
			}),
			widget.NewLabel("ip address"),
			widget.NewSeparator(),
			widget.NewEntry(),
			widget.NewLabel("port"),
			widget.NewSeparator(),
			widget.NewEntry(),
			button,
		),
		func(b bool) {
			println("test")
		},
		window)
	dl.Resize(fyne.NewSquareSize(400))
	dl.Show()
}
