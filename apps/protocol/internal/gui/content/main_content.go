package content

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"go.uber.org/fx"
)

type MainContentDependency struct {
	fx.In
	ConnectionManager *fyne.Container `name:"connectionManager"`
	ContentTab        *container.AppTabs
}

func NewMainContent(dependency MainContentDependency) *container.Split {
	hsplit := container.NewHSplit(
		dependency.ConnectionManager,
		dependency.ContentTab,
	)
	hsplit.SetOffset(0.2)
	return hsplit
}
