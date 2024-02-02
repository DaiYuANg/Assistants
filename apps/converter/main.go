package main

import (
	"context"
	fyneApp "fyne.io/fyne/v2/app"
	"io"
	"log"
	"os"

	_ "github.com/u2takey/ffmpeg-go"
	"github.com/wader/goutubedl"
)

func main() {
	app := fyneApp.NewWithID("converter")
	result, err := goutubedl.New(context.Background(), "https://www.youtube.com/watch?v=jgVhBThJdXc", goutubedl.Options{})
	if err != nil {
		log.Fatal(err)
	}
	downloadResult, err := result.Download(context.Background(), "best")
	if err != nil {
		log.Fatal(err)
	}
	defer downloadResult.Close()
	f, err := os.Create("output")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	io.Copy(f, downloadResult)
	app.Run()
}
