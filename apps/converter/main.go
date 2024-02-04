package main

import (
	"fmt"
	"fyne.io/fyne/v2"
	fyneApp "fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/widget"
	_ "github.com/u2takey/ffmpeg-go"
)

func main() {
	app := fyneApp.NewWithID("converter")
	window := app.NewWindow("co")
	window.SetContent(container.NewCenter(
		widget.NewButton("Choose File", func() {
			dialog.NewFileOpen(func(closer fyne.URIReadCloser, err error) {
				fmt.Println(closer)
			}, window).Show()
		}),
	))
	window.CenterOnScreen()
	window.Resize(fyne.NewSize(300, 300))
	window.Show()
	//result, err := goutubedl.New(context.Background(), "https://www.youtube.com/watch?v=jgVhBThJdXc", goutubedl.Options{})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//downloadResult, err := result.Download(context.Background(), "best")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer downloadResult.Close()
	//f, err := os.Create("output")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer f.Close()
	//io.Copy(f, downloadResult)
	app.Run()
}
