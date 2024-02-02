package component

import "fyne.io/fyne/v2/widget"

func NewButton(label string, tapped func(button *widget.Button)) *widget.Button {
	var button *widget.Button
	var callback = func() {
		tapped(button)
	}
	button = widget.NewButton(label, callback)
	return button
}
