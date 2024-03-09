package content

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/widget"
	widgetx "fyne.io/x/fyne/widget"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"protocol/internal/model"
)

type ConnectionManagerDependency struct {
	fx.In
	DB     *gorm.DB
	Logger *zap.Logger
}

func NewConnectionManager(dependency ConnectionManagerDependency) *fyne.Container {
	var connections []model.Connection
	dependency.DB.Find(&connections)
	widgetx.NewFileTree(storage.NewFileURI("~"))
	list := widget.NewList(func() int {
		return 10
	}, func() fyne.CanvasObject {
		return container.NewHBox(
			widget.NewLabel("test"),
		)
	}, func(id widget.ListItemID, object fyne.CanvasObject) {
		fmt.Println(id)
	})
	list.OnSelected = func(id widget.ListItemID) {
		fmt.Println(id)
	}
	tree := widget.NewAccordion(
		widget.NewAccordionItem("1", list),
	)
	vlayout := layout.NewVBoxLayout()
	scroll := container.NewVScroll(
		tree,
	)
	scroll.SetMinSize(fyne.NewSize(30, 600))

	manager := container.New(
		vlayout,
		widget.NewEntry(),
		scroll,
	)
	return manager
}
