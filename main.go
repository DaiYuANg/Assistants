package main

import (
	"github.com/samber/lo"
	"protocol/assistant/cmd"
)

func main() {
	lo.Must0(cmd.Execute())
}
