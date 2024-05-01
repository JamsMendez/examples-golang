package file

import (
	"log"
	"os"
)

func DeleteFile() {
	err := os.Remove("file.txt")
	if err != nil {
		log.Fatal(err)
	}
}
