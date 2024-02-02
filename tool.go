//go:build tools
// +build tools

package main

import (
	_ "fyne.io/fyne/v2/cmd/fyne"
	_ "github.com/cosmtrek/air"
	_ "github.com/goreleaser/goreleaser"
	_ "golang.org/x/tools/cmd/stringer"
)
