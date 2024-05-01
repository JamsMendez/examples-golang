package file

import (
	"fmt"
	"log"
	"os"
)

func Info() {
	info, err := os.Stat("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(info)
}
