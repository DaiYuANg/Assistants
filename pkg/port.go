package pkg

import (
	"log/slog"
	"net"
)

func FindAvailablePort() (int, error) {
	listener, err := net.Listen("tcp", ":0")
	if err != nil {
		return 0, err
	}
	defer func(listener net.Listener) {
		err := listener.Close()
		if err != nil {
			slog.Error(err.Error())
		}
	}(listener)

	address := listener.Addr().(*net.TCPAddr)
	return address.Port, nil
}
