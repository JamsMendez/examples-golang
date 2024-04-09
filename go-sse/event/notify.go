package event

import (
	"fmt"
	"io"
	"sync"
)

type HandlerEvent interface {
	Subscribe(clientID string) (*client, error)
	Unsubscribe(clientID string)
	Emit(ClientID string, msg Message, w io.Writer)
	Broadcast(msg Message)
}

type Message struct {
	Name string
	Data any
}

type handlerEvent struct {
	clients map[string]*client
	sync.Mutex
}

func NewHandlerEvent() *handlerEvent {
	return &handlerEvent{
		clients: make(map[string]*client),
	}
}

func (h *handlerEvent) Subscribe(clientID string) (*client, error) {
	h.Lock()
	defer h.Unlock()

	if _, ok := h.clients[clientID]; ok {
		return nil, fmt.Errorf("client %s already subscribed", clientID)
	}

	c := newClient(clientID)

	h.clients[clientID] = c

	return c, nil
}

func (h *handlerEvent) Unsubscribe(clientID string) {
	h.Lock()
	defer h.Unlock()

	delete(h.clients, clientID)
}

func (h *handlerEvent) Emit(clientID string, msg Message, w io.Writer) {
	h.Lock()
	defer h.Unlock()

	client, ok := h.clients[clientID]
	if !ok {
		return
	}

	client.Send(msg, w)
}

func (h *handlerEvent) Broadcast(msg Message) {
	h.Lock()
	defer h.Unlock()
	for _, client := range h.clients {
		client.notify(msg)
	}
}
