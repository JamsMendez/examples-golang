package main

import "fmt"

func main() {
	runRWMutex(false)
}

func logln(debug bool, args ...any) {
    if  !debug {
        return
    }

    fmt.Println(args...)
}
