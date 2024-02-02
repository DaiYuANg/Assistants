package ui

import (
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"gorm.io/gorm"
	"log"
	"protocol/assistant/internal/model"
	"protocol/assistant/internal/protocol"
)

func NewToolBar(db *gorm.DB) *widget.Toolbar {
	toolbar := widget.NewToolbar(
		widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
			log.Println("New document")
			connection := model.Connection{Protocol: protocol.TCP_SERVER}
			db.Create(&connection)
		}),
		widget.NewToolbarSeparator(),
		widget.NewToolbarAction(theme.HistoryIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentCopyIcon(), func() {}),
		widget.NewToolbarAction(theme.ContentPasteIcon(), func() {}),
		widget.NewToolbarSpacer(),
		widget.NewToolbarAction(theme.HelpIcon(), func() {
			log.Println("Display help")
		}))
	return toolbar
}
