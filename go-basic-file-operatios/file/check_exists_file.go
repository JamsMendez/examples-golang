package file

import (
	"log"
	"os"
)

func CheckFile() {
	info, err := os.Stat("file.txt")
	if err != nil {
		if os.IsNotExist(err) {
			log.Fatal("file does not exist")
		}
	}

	log.Println(info)
}
