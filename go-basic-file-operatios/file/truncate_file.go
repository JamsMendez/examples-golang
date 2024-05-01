package file

import (
	"log"
	"os"
)

func TruncateFile() {
	err := os.Truncate("file.txt", 100)
	if err != nil {
		log.Fatal(err)
	}
}
