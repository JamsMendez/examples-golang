package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func MyGreeterHandler(w http.ResponseWriter,r *http.Request) {
  Greet(w, "word")
}

func  Greet(writer io.Writer, name string) {
  fmt.Fprintf(writer, "Hello %s", name)
}

func main() {
  log.Fatal(http.ListenAndServe(":5000", http.HandlerFunc(MyGreeterHandler)))
}
