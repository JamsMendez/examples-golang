package main

import (
	"fmt"
	"io"
	"os"
)

func deferRunOrder() {
  // The defer statements run in reverse order, from last to first
  for i := 0; i <= 4; i++ {
    defer fmt.Println("Deferred ", -i)
    fmt.Println("Regular", i)
  }
}

func deferCloseFile() {
  newFile, err := os.Create("learnGo.txt")
  if err != nil {
    fmt.Println("Error: Could not create file.")
    return
  }

  defer newFile.Close()

  if _, err = io.WriteString(newFile, "Learning Go!"); err != nil {
    fmt.Println("Error: Could not write to file.")
    return
  }

  newFile.Sync()
}

func deferVariables() {
  var message = "Hola"

  // message has value "Hola"
  defer fmt.Println("Message.2: ", message)

  // message has value "Mundo"
  defer func () {
    fmt.Println("Message.1: ", message)
  }()


  message = "Mundo"

  fmt.Println("Finish...")
}
