package server

import (
	"log"
	"net/http"

	"go-sse/event"
	"go-sse/handler"
)

func Start() {
	server := http.NewServeMux()

	handlerEvent := event.NewHandlerEvent()

	server.HandleFunc("/", handler.Home)
	server.HandleFunc("/notify", handler.Notify(handlerEvent))
	server.HandleFunc("/test", handler.EmitTestEvent(handlerEvent))
	server.HandleFunc("/broadcast", handler.EmitNewMessage(handlerEvent))

	log.Println("Server started at port 3000")

	err := http.ListenAndServe(":3000", server)
	if err != nil {
		log.Fatal(err)
	}
}
