package dialog

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"go.uber.org/fx"
)

type ConnectionDialogDependency struct {
	fx.In
	Window fyne.Window
}
type CreateConnectionDialog struct {
	dialog.CustomDialog
	Result ConnectionDialogResult
}

func (d CreateConnectionDialog) ClearAndShow() {
	d.Show()
}

func NewCreateConnectionDialog(dependency ConnectionDialogDependency) *dialog.CustomDialog {
	var createConnectionDialog *dialog.CustomDialog
	var formContainer, _, _ = NewCreateConnectionForm(ConnectionFormParameter{
		OnCancel: func() {
			createConnectionDialog.Hide()
		},
	})

	createConnectionDialog = dialog.NewCustomWithoutButtons(
		"Create connection",
		container.NewVBox(
			formContainer,
		),
		dependency.Window,
	)
	createConnectionDialog.Resize(fyne.NewSize(300, 250))

	return createConnectionDialog
}
