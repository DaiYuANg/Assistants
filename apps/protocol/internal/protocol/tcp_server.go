package protocol

import (
	"assistants/pkg"
	"fmt"
	"net"
	"strconv"
)

func NewTcpServer() {
	port, err := pkg.FindAvailablePort()
	if err != nil {
		return
	}
	address := "localhost:" + strconv.Itoa(port)
	// 监听地址和端口
	listener, err := net.Listen("tcp", address)
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			panic(err)
		}
	}(listener)

	fmt.Println("Server listening on " + address)

	// 无限循环等待客户端连接
	for {
		// 等待客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		// 启动一个新的 goroutine 处理连接
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// 在连接关闭之前一直处理客户端的请求
	defer func(conn net.Conn) {
		err := conn.Close()
		if err != nil {

		}
	}(conn)

	// 处理客户端的请求
	// 这里可以编写处理逻辑，例如读取/写入数据
	fmt.Println("Client connected:", conn.RemoteAddr())

	// 读取客户端发送的数据
	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}

	// 处理接收到的数据
	fmt.Println("Received data:", string(buffer))
}
