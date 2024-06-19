package main

import (
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

const port = 3000

func main() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	number := r.Intn(100)

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, _ = fmt.Fprintf(w, "Service N %d - %s", number, time.Now().Format(time.DateTime))
	})

	mux.HandleFunc("/exit", func(w http.ResponseWriter, r *http.Request) {
		os.Exit(1)
	})

	fmt.Printf("Server http running in port %d\n", port)

	addr := fmt.Sprintf(":%d", port)
	err := http.ListenAndServe(addr, mux)
	log.Fatal(err)
}
