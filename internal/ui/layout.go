package ui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/binding"
	"fyne.io/fyne/v2/widget"
	"log"
	"protocol/assistant/internal/socket"
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
		socket.NewSocketServer()
	}()
}

func NewLayout() *fyne.Container {
	text, _ := boundString.Get()
	var button = widget.NewButton(text, onTap)
	boundString.AddListener(binding.NewDataListener(func() {
		println("change")
		var te, _ = boundString.Get()
		button.SetText(te)
	}))
	mainLayout := container.NewBorder(
		// top
		container.NewVBox(
		//widget.NewLabel("Top"),
		),
		// bottom
		container.NewHBox(
			widget.NewLabel("Left 1"),
			widget.NewLabel("Left 2"),
			widget.NewProgressBarInfinite(),
		),
		// left
		container.NewVBox(
			widget.NewLabel("ip address"),
			widget.NewEntry(),
			widget.NewSeparator(),
			widget.NewLabel("protocol"),
			widget.NewSelect(options, func(selected string) {
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
		container.NewVBox(
		//widget.NewLabel("Right 1"),
		//widget.NewLabel("Right 2"),
		),
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
