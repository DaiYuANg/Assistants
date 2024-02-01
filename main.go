package main

import (
	"github.com/samber/lo"
	"protocol/assistant/cmd"
)

func main() {
	//a := app.New()
	//w := a.NewWindow("Hello")
	//hello := widget.NewLabel("Hello Fyne!")
	//w.SetContent(container.NewVBox(
	//	hello,
	//	widget.NewButton("Hi!", func() {
	//		hello.SetText("Welcome :)")
	//	}),
	//))
	//
	//w.ShowAndRun()
	lo.Must0(cmd.Execute())
}
