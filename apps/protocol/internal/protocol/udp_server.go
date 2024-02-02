package protocol

import (
	"assistants/pkg"
	"fmt"
	"net"
	"strconv"
)

func NewUdpServer() {
	port, err := pkg.FindAvailablePort()
	if err != nil {
		return
	}
	address := "localhost:" + strconv.Itoa(port)
	addr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer conn.Close()

	fmt.Println("UDP Server is listening on", addr)

	// 处理UDP请求
	for {
		handleUDPRequest(conn)
	}
}

func handleUDPRequest(conn *net.UDPConn) {
	buffer := make([]byte, 1024)

	// 读取客户端发送的数据
	n, addr, err := conn.ReadFromUDP(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	receivedData := string(buffer[:n])
	fmt.Printf("Received data from %s: %s\n", addr, receivedData)

	// 向客户端发送响应数据
	response := "Hello, client!"
	_, err = conn.WriteToUDP([]byte(response), addr)
	if err != nil {
		fmt.Println("Error sending data:", err)
		return
	}
}
