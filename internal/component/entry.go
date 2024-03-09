package component

import (
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
)

type EntryOption struct {
	Data     binding.String
	OnChange func(s string)
}

func NewEntry(option EntryOption) {
	var entry = widget.NewEntryWithData(option.Data)
	entry.OnChanged = option.OnChange
}
