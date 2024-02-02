package global_toolbar

import (
	"fyne.io/fyne/v2/widget"
)

func NewToolBar(actions []widget.ToolbarItem) *widget.Toolbar {
	return widget.NewToolbar(actions...)
}
