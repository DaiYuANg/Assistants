package protocol

type (
	Protocol string
	Base     uint8
)

const (
	TcpServer Protocol = "TCP_SERVER"
	UdpServer Protocol = "UDP_SERVER"
	TcpClient Protocol = "TCP_CLIENT"
	UdpClient Protocol = "UDP_CLIENT"
	BASE2              = 2
	BASE8              = 8
	BASE10             = 10
	BASE16             = 16
	SIZE               = 5
)

var (
	AvailableProtocols = []string{
		string(TcpServer),
		string(UdpServer),
		string(TcpClient),
		string(UdpClient),
	}
)

type CreateConnectionOption struct {
	Protocol Protocol
	Address  string
	Port     string
}

type AcceptorOption struct {
	SaveToFile     string
	Wrap           bool
	ShowAcceptTime bool
}

type SenderOption struct {
	Base         Base
	ScheduleTime uint
}
