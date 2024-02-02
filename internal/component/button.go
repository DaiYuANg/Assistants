package component

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

func NewButton(label string, tapped func(button *widget.Button)) *widget.Button {
	var button *widget.Button
	var callback = func() {
		tapped(button)
	}
	button = widget.NewButton(label, callback)
	return button
}

func NewButtonLabelBind(data binding.String, tapped func(button *widget.Button)) *widget.Button {
	var button *widget.Button
	data.AddListener(binding.NewDataListener(func() {
		println("change")
		var te, _ = data.Get()
		button.SetText(te)
	}))
	button = NewButton("", tapped)
	return button
}
