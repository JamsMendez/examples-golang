package file

import (
	"log"
	"os"
)

func RenameMoveFile() {
	oldPath := "filename.txt"
	newPath := "file.txt"

	err := os.Rename(oldPath, newPath)
	if err != nil {
		log.Fatal(err)
	}
}
