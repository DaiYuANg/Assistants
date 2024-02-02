package global_toolbar

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type TBDependency struct {
	fx.In
	DB                     *gorm.DB
	Window                 fyne.Window
	CreateConnectionDialog *dialog.CustomDialog
}

func NewCreateConnectionAction(dependency TBDependency) widget.ToolbarItem {
	ta := widget.NewToolbarAction(theme.DocumentCreateIcon(), func() {
		//pdialog.NewCreateConnectionDialog(dependency.Window)
		dependency.CreateConnectionDialog.Show()
	})
	return ta
}
