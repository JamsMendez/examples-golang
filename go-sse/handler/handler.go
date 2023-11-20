package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go-sse/event"
)

func Home(w http.ResponseWriter, r *http.Request) {
	http.FileServer(http.Dir("./static/home")).ServeHTTP(w, r)
}

func Notify(handlerEvent event.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		flusher, ok := w.(http.Flusher)
		if !ok {
			http.Error(w, "Streaming unsupported!", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		w.Header().Set("Content-Type", "text/event-stream")
		w.Header().Set("Cache-Control", "no-cache")
		w.Header().Set("Connection", "keep-alive")

		clientID := r.RemoteAddr + fmt.Sprintf("%d", time.Now().Unix())

		client, err := handlerEvent.Subscribe(clientID)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		fmt.Println("New client subscribed:", clientID)

		defer handlerEvent.Unsubscribe(clientID)

		for {
			fmt.Println("Waiting for message...")
			select {
			case <-r.Context().Done():
				fmt.Println("client Unsubscribe:", clientID)
				return

			case msg := <-client.Receive():
				fmt.Println("Client sending message...")

				client.Send(msg, w)

				flusher.Flush()

				fmt.Println("Client message send!")
			}
		}
	}
}

func EmitTestEvent(handlerEvent event.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, _ *http.Request) {
		data := map[string]string{
			"message": fmt.Sprintf("Hello, world! %d", time.Now().Unix()),
		}

		msg := event.Message{
			Name: "new_message",
			Data: data,
		}

		fmt.Println("Handler sending broadcast...")

		handlerEvent.Broadcast(msg)

		w.WriteHeader(http.StatusOK)

		fmt.Println("Handler broadcast sent!")
	}
}

func EmitNewMessage(handlerEvent event.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data map[string]any

		err := json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		msg := event.Message{
			Name: "delete_message",
			Data: data,
		}

		fmt.Println("Sending broadcast ...")

		handlerEvent.Broadcast(msg)

		w.WriteHeader(http.StatusOK)

		fmt.Println("Broadcast sent!")
	}
}
