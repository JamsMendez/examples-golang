package main

import "fmt"

func main() {
	core := &Core{}
	logger := &LoggerPlugin{}

	core.Register("logger", logger)
	core.Active("logger")

	fmt.Println(core.Execute("logger"))
	fmt.Println(core.Execute("metrics"))
}
