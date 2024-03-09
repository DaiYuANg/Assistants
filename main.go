package main

import (
	_ "github.com/joho/godotenv/autoload"
	"github.com/samber/lo"
	"protocol/cmd"
)

func main() {
	lo.Must0(cmd.Execute())
}
