package cmd

import (
	"asssistant/protocol/internal/context"
	"fmt"
	"go.uber.org/fx"
)

var contextModule = fx.Module("contextModule", fx.Provide(newContext), fx.Invoke(startContext))

func newContext() *context.ProtocolContext {
	c := context.NewProtocolContext()
	return c
}

func startContext(pctx *context.ProtocolContext) {
	println("echo")
	go func() {
		for {
			println("on whil")
			select {
			case message := <-pctx.MessageChan:
				println("message")
				fmt.Println(message)
			case <-pctx.StopChan:
				break
			}
		}
	}()
}
