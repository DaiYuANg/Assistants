package cmd

import (
	"fmt"
	"go.uber.org/fx"
	"net"
	"os"
)

var infoModule = fx.Module("info", fx.Provide(NewEnvironment))

type Environment struct {
	IPAddress []string
}

func NewEnvironment() *Environment {
	name, err := os.Hostname()
	addrs, err := net.LookupHost(name)
	if err != nil {
		fmt.Printf("Oops: %v\n", err)
		return nil
	}
	return &Environment{IPAddress: addrs}
}
