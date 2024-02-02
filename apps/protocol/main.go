package main

import (
	"asssistant/protocol/cmd"
	"github.com/samber/lo"
)

func main() {
	lo.Must0(cmd.Execute())
}
