package protocol

import (
	"fmt"
	"net"
)

func NewTcpClient(address string) {
	conn, err := net.Dial("tcp", address)
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()
}
