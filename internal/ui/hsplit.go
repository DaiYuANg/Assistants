package ui

import (
	"fyne.io/fyne/v2/widget"
	"go.uber.org/fx"
)

type MainContentParam struct {
	fx.In
	ConnectionTree *widget.Tree `name:"connectionTree"`
}

func NewMainContent(param MainContentParam) {

}
