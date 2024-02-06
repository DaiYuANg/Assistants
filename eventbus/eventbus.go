package eventbus

import "sync"

type Message[T any] struct {
	Content T
	ReplyTo chan T // 回复通道
}

type EventBus[T any] struct {
	subscribers map[string][]chan Message[T]
	mu          sync.Mutex
}

// NewEventBus 创建一个新的事件总线
func NewEventBus[T any]() *EventBus[T] {
	return &EventBus[T]{
		subscribers: make(map[string][]chan Message[T]),
	}
}

// Subscribe 订阅指定主题的事件
func (bus *EventBus[T]) Subscribe(topic string) chan Message[T] {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	ch := make(chan Message[T])
	bus.subscribers[topic] = append(bus.subscribers[topic], ch)

	return ch
}

// Publish 发布事件到指定主题
func (bus *EventBus[T]) Publish(topic string, message Message[T]) {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	if subscribers, ok := bus.subscribers[topic]; ok {
		for _, ch := range subscribers {
			go func(ch chan Message[T]) {
				ch <- message
			}(ch)
		}
	}
}

// Unsubscribe 取消订阅指定主题的事件
func (bus *EventBus[T]) Unsubscribe(topic string, ch chan Message[T]) {
	bus.mu.Lock()
	defer bus.mu.Unlock()

	if subscribers, ok := bus.subscribers[topic]; ok {
		var updatedSubscribers []chan Message[T]
		for _, subscriber := range subscribers {
			if subscriber != ch {
				updatedSubscribers = append(updatedSubscribers, subscriber)
			}
		}
		bus.subscribers[topic] = updatedSubscribers
		close(ch)
	}
}
