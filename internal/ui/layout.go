package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	xwidget "fyne.io/x/fyne/widget"
	"go.uber.org/fx"
	"log"
	"protocol/assistant/internal/protocol"
)

var options = []string{"Option 1", "Option 2", "Option 3"}
var boundString = binding.NewString()

func init() {
	err := boundString.Set("connect")
	if err != nil {
		panic(err)
	}
}

func onTap() {
	err := boundString.Set("disConnect")
	if err != nil {
		panic(err)
	}
	go func() {
		protocol.NewTcpServer()
	}()
}

type LayoutParam struct {
	fx.In
	window  fyne.Window
	toolbar *widget.Toolbar
}

func NewLayout(window fyne.Window, toolbar *widget.Toolbar) *fyne.Container {
	text, _ := boundString.Get()
	var button = widget.NewButton(text, onTap)
	boundString.AddListener(binding.NewDataListener(func() {
		println("change")
		var te, _ = boundString.Get()
		button.SetText(te)
	}))
	mainLayout :=
		container.NewBorder(
			// top
			toolbar,
			// bottom
			container.NewHBox(
				widget.NewLabel("Left 1"),
				widget.NewLabel("Left 2"),
				widget.NewProgressBarInfinite(),
			),
			// left
			container.NewVBox(
				widget.NewLabel("ip address"),
				xwidget.NewCompletionEntry([]string{}),
				widget.NewSeparator(),
				widget.NewLabel("protocol"),
				widget.NewSelect(options, func(selected string) {
					NewCreateConnectionDialog(window)
					fmt.Println(selected)
				}),
				widget.NewSeparator(),
				button,
				widget.NewLabel("Bottom"),
				&widget.Form{
					Items: []*widget.FormItem{
						{Text: "Entry", Widget: widget.NewEntry()}},
					OnSubmit: func() {
						log.Println("Form submitted:", widget.NewEntry().Text)
						log.Println("multiline:", widget.NewMultiLineEntry().Text)
					},
				},
			),
			// right
			nil,
			// center
			container.NewVBox(
				container.NewAppTabs(
					container.NewTabItem("Tab 1", container.NewVBox(
						widget.NewLabel("Center 1"),
						NewMessageContent(),
					)),
					container.NewTabItem("Tab 2", widget.NewLabel("World!")),
				),
			),
		)
	return mainLayout
}

func NewMessageContent() *fyne.Container {
	text := widget.NewRichTextWithText("test")
	text.Resize(fyne.NewSize(50, 100))
	scroll := container.NewScroll(
		text,
	)
	scroll.Resize(fyne.NewSize(500, 500))
	return container.NewVBox(
		scroll,
	)
}
