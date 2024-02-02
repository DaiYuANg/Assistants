package context

type ProtocolContext struct {
	MessageChan chan struct{}
	StopChan    chan struct{}
}

func NewProtocolContext() *ProtocolContext {
	return &ProtocolContext{
		MessageChan: make(chan struct{}),
		StopChan:    make(chan struct{}),
	}
}
