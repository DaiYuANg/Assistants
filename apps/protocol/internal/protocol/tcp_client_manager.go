package protocol

import (
	"fmt"
	"github.com/google/uuid"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"net"
)

type TcpClientManager struct {
	Connections map[string]net.Conn
}

type TcpClientManagerDependency struct {
	fx.In
	Logger *zap.Logger
}

func NewTcpClientManager(dependency TcpClientManagerDependency) *TcpClientManager {
	fmt.Println("123")
	manager := &TcpClientManager{
		Connections: make(map[string]net.Conn),
	}

	return manager
}

func (m *TcpClientManager) Create(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting:", err)
	}
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {
			panic(err)
		}
	}(conn)
	newUUID := uuid.New()
	m.Connections[newUUID.String()] = conn
}
