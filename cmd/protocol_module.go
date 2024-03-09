package cmd

import (
	"fmt"
	"github.com/gookit/event"
	"go.uber.org/fx"
	"protocol/internal/protocol"
)

var protocolModule = fx.Module("protocolModule",
	fx.Provide(protocol.NewTcpClientManager),
	fx.Invoke(tcpClientManagerListener),
)

func tcpClientManagerListener(manager *protocol.TcpClientManager) {
	event.On("evt1", event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("handle event: %s\n", e.Name())
		manager.Create("")
		return nil
	}), event.Normal)
}
