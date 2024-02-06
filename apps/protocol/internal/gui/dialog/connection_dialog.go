package dialog

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"github.com/gookit/event"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type ConnectionDialogDependency struct {
	fx.In
	Window fyne.Window
	Logger *zap.Logger
}
type CreateConnectionDialog struct {
	dialog.CustomDialog
	Result ConnectionDialogResult
}

func (d CreateConnectionDialog) ClearAndShow() {
	d.Show()
}

func NewCreateConnectionDialog(dependency ConnectionDialogDependency) *dialog.CustomDialog {
	fmt.Println(321)
	var createConnectionDialog *dialog.CustomDialog
	var formContainer, _, _ = NewCreateConnectionForm(ConnectionFormParameter{
		OnCancel: func() {
			createConnectionDialog.Hide()
		},
		OnSave: func(result ConnectionDialogResult) {
			event.MustFire("evt1", event.M{"arg0": "val0", "arg1": "val1"})
			dependency.Logger.Log(zap.InfoLevel, "create connection", zap.String("create", "c"))
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
