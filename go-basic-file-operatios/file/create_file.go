package file

import (
	"log"
	"os"
)

func CreateFile() {
	file, err := os.Create("file.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = file.Close()
		if err != nil {
			log.Println("file close error: ", err)
		}
	}()

	log.Println(file)
}
