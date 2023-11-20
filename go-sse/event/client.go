package event

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
)

const format = "event:%s\ndata:%s\n\n"

type client struct {
	id      string
	message chan Message
}

func newClient(id string) *client {
	return &client{
		id:      id,
		message: make(chan Message),
	}
}

func (c *client) Receive() <-chan Message {
	return c.message
}

func (c *client) Send(msg Message, w io.Writer) {
	buffer, err := json.Marshal(msg.Data)
	if err != nil {
		log.Println(err)
	}

	data := string(buffer)
	_, err = fmt.Fprintf(w, format, msg.Name, data)
	if err != nil {
		log.Println(err)
	}
}

func (c *client) notify(msg Message) {
	c.message <- msg
}
