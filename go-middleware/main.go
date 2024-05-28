package main

import (
	"log"
	"net/http"

	"go-middleware/handler"
)

func main() {
	mux := http.NewServeMux()

	h := handler.NewHandler()

	mux.HandleFunc("GET /user/list", h.ListUser)

	muxRedirect := handler.TrailingSlashRedirect(mux)
	muxAuth := handler.IsAuth(muxRedirect)
	muxID := handler.RequestID(muxAuth)
	muxServer := handler.ResponseServer(muxID, "JamsMendez")
	muxLog := handler.LoggingMiddleware(muxServer)

	err := http.ListenAndServe(":3000", muxLog)
	log.Fatal(err)
}
