package protocol

type (
	Protocol string
	Base     uint8
)

const (
	TCP_SERVER Protocol = "TCP_SERVER"
	UDP_SERVER Protocol = "UDP_SERVER"
	BASE2               = 2
	BASE8               = 8
	BASE10              = 10
	BASE16              = 16
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
